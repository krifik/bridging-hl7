package entity

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	// gorm.Model
	// Name string
	// Type string
	ID             int          `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	ProductImageId int          `gorm:"unique"`
	ProductImage   ProductImage `gorm:"foreignKey:ID;references:ProductImageId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryId     int          `gorm:"unique"`
	Category       Category     `gorm:"foreignKey:CategoryId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PartnerId      int          `gorm:"unique"`
	Partner        Partner      `gorm:"foreignKey:PartnerId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name           string       `gorm:"size:256;not null"`
	Desc           string       `gorm:"size:256;not null"`
	Stock          int
	Price          decimal.Decimal `gorm:"type:decimal(16,2)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
