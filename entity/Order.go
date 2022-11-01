package entity

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	ID        int `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	PaymentId int `gorm:"index"`
	DriverId  int `gorm:"index"`
	ProductId int `gorm:"index"`
	UserId    int `gorm:"index"`
	User      User
	Payment   Payment
	Driver    Driver
	// Product   Product `gorm:"foreignKey:ProductId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TaxAmount       decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxPercent      decimal.Decimal `gorm:"type:decimal(16,2)"`
	DiscountAmount  decimal.Decimal `gorm:"type:decimal(16,2)"`
	DiscountPercent decimal.Decimal `gorm:"type:decimal(16,2)"`
	Total           decimal.Decimal `gorm:"type:decimal(16,2)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
