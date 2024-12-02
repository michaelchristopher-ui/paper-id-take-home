package disbursement

import (
	"context"
	"errors"
	"fmt"
	"paperid-entry-task/internal/pkg/adapters/accountrepo"
	"paperid-entry-task/internal/pkg/adapters/disbursementsvc"
	"paperid-entry-task/internal/pkg/adapters/journalrepo"
	"time"
)

var _ disbursementsvc.Adapter = &Disbursement{}

type Disbursement struct {
	journalRepo journalrepo.RepoAdapter
	accountRepo accountrepo.RepoAdapter
}

func New(journalRepo journalrepo.RepoAdapter, accountRepo accountrepo.RepoAdapter) disbursementsvc.Adapter {
	return &Disbursement{
		journalRepo: journalRepo,
		accountRepo: accountRepo,
	}
}

func (d *Disbursement) Disburse(ctx context.Context, req disbursementsvc.DisburseReq) (err error) {
	tx := d.journalRepo.CreateTx()
	if tx == nil {
		return errors.New("[Disburse] error when creating transaction on disburse")
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	now := time.Now()

	entries := GenerateDebitAndCreditForDisbursement(req.AccountIdFrom, req.AccountIdTo, now, req.Amount, req.Description)
	for _, entry := range entries {
		err = d.journalRepo.Insert(ctx, tx, entry)
		if err != nil {
			return fmt.Errorf("[Disburse] error when inserting: %s", err.Error())
		}
	}

	//Actual transfer
	d.accountRepo.UpdateBalance(ctx, tx, accountrepo.UpdateBalance{
		ID:          req.AccountIdFrom,
		BalanceIncr: -req.Amount,
	})
	d.accountRepo.UpdateBalance(ctx, tx, accountrepo.UpdateBalance{
		ID:          req.AccountIdTo,
		BalanceIncr: req.Amount,
	})
	//End Actual Transfer

	res := tx.Commit()
	if res.Error != nil {
		return fmt.Errorf("[Disburse] error when inserting: %s", err.Error())
	}
	return nil
}
