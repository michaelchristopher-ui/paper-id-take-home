package model

import "time"

type JournalEntry struct {
	ID           uint `gorm:"primaryKey"`
	EntryDate    time.Time
	Description  string
	AccountID    int
	DebitAmount  float64
	CreditAmount float64
	Account      Account `gorm:"foreignKey:AccountID"`
}
