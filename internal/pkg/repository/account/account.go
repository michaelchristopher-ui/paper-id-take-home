package account

import (
	"context"
	"errors"
	"paperid-entry-task/internal/data/model"
	"paperid-entry-task/internal/pkg/adapters/accountrepo"
	"time"

	"gorm.io/gorm"
)

var _ accountrepo.RepoAdapter = &AccountRepo{}

type AccountRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) accountrepo.RepoAdapter {
	return &AccountRepo{
		db: db,
	}
}

func (a *AccountRepo) Create(ctx context.Context, db *gorm.DB, req accountrepo.Create) error {
	if db == nil {
		db = a.db
	}
	now := time.Now()
	res := db.Create(&model.Account{
		AccountName: req.AccountName,
		Balance:     req.Balance,
		CreatedAt:   now,
		UpdatedAt:   now,
		IsActive:    req.IsActive,
	})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}

func (a *AccountRepo) UpdateBalance(ctx context.Context, db *gorm.DB, req accountrepo.UpdateBalance) error {
	if db == nil {
		db = a.db
	}

	account := &model.Account{}

	// Check if the account exists
	if err := db.WithContext(ctx).Where(&model.Account{ID: req.ID}).Find(account).Error; err != nil {
		return err
	}

	// Update the balance
	if account.Balance+req.BalanceIncr < 0 {
		return errors.New("insufficient balance")
	}

	account.UpdatedAt = time.Now()
	account.Balance += req.BalanceIncr

	res := db.Save(account)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no rows affected") // Consider a more specific error
	}
	return nil
}
