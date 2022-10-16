package entity

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	Id        int     `gorm:"primaryKey,not null, autoIncrement;uniqueIndex"`
	PaymentId int     `gorm:"unique"`
	DriverId  int     `gorm:"unique"`
	ProductId int     `gorm:"unique"`
	UserId    int     `gorm:"unique"`
	User      User    `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Payment   Payment `gorm:"foreignKey:Id;references:PaymentId"`
	Driver    Driver  `gorm:"foreignKey:Id;references:DriverId"`
	Product   Product `gorm:"foreignKey:Id;references:ProductId"`

	TaxAmount       decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxPercent      decimal.Decimal `gorm:"type:decimal(16,2)"`
	DiscountAmount  decimal.Decimal `gorm:"type:decimal(16,2)"`
	DiscountPercent decimal.Decimal `gorm:"type:decimal(16,2)"`
	Total           decimal.Decimal `gorm:"type:decimal(16,2)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
