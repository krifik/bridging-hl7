package entity

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	ID        int     `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	PaymentId int     `gorm:"unique"`
	DriverId  int     `gorm:"unique"`
	ProductId int     `gorm:"unique"`
	UserId    int     `gorm:"unique"`
	User      User    `gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payment   Payment `gorm:"foreignKey:ID;references:PaymentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Driver    Driver  `gorm:"foreignKey:ID;references:DriverId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Product   Product `gorm:"foreignKey:ID;references:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	TaxAmount       decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxPercent      decimal.Decimal `gorm:"type:decimal(16,2)"`
	DiscountAmount  decimal.Decimal `gorm:"type:decimal(16,2)"`
	DiscountPercent decimal.Decimal `gorm:"type:decimal(16,2)"`
	Total           decimal.Decimal `gorm:"type:decimal(16,2)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
