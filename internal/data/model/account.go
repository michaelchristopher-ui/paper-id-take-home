package model

import (
	"time"
)

type Account struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	AccountName string
	Balance     float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActive    bool
}
