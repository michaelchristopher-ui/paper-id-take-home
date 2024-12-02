package accountrepo

import (
	"context"

	"gorm.io/gorm"
)

type RepoAdapter interface {
	Create(ctx context.Context, db *gorm.DB, req Create) error
	UpdateBalance(ctx context.Context, db *gorm.DB, req UpdateBalance) error
}

type Create struct {
	ID          int
	AccountName string
	Balance     float64
	IsActive    bool
}

type UpdateBalance struct {
	ID          int
	BalanceIncr float64
}
