package entity

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Payment struct {
	Id        int             `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	OrderId   int             `gorm:"index"`
	Amount    decimal.Decimal `gorm:"type:decimal(16,2)"`
	Method    string          `gorm:"size:100"`
	Status    string          `gorm:"size:100"`
	Token     string          `gorm:"size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
