package journalrepo

import (
	"context"
	"paperid-entry-task/internal/data/model"

	"gorm.io/gorm"
)

type RepoAdapter interface {
	Insert(ctx context.Context, db *gorm.DB, journal model.JournalEntry) error
	CreateTx() *gorm.DB
}
