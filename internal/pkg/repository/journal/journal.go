package journal

import (
	"context"
	"errors"
	"paperid-entry-task/internal/data/model"
	"paperid-entry-task/internal/pkg/adapters/journalrepo"

	"gorm.io/gorm"
)

var _ journalrepo.RepoAdapter = &JournalRepo{}

type JournalRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) journalrepo.RepoAdapter {
	return &JournalRepo{
		db: db,
	}
}

func (j *JournalRepo) CreateTx() *gorm.DB {
	return j.db.Begin()
}

func (j *JournalRepo) Insert(ctx context.Context, db *gorm.DB, journal model.JournalEntry) error {
	if db == nil {
		db = j.db
	}
	res := db.Create(&journal)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return nil
}
