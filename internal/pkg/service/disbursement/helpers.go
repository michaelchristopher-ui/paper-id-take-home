package disbursement

import (
	"fmt"
	"paperid-entry-task/internal/data/model"
	"time"
)

func GenerateDebitAndCreditForDisbursement(from int, to int, entry time.Time, amount float64, description string) []model.JournalEntry {
	return []model.JournalEntry{
		{
			EntryDate:   entry,
			Description: fmt.Sprintf("DEBIT: %s", description),
			AccountID:   from,
			DebitAmount: amount,
		},
		{
			EntryDate:    entry,
			Description:  fmt.Sprintf("CREDIT: %s", description),
			AccountID:    to,
			CreditAmount: amount,
		},
	}
}
