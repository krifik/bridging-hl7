package entity

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID             int `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	CategoryId     int
	Category       Category
	PartnerId      int
	Partner        Partner
	ProductImageId int
	ProductImage   ProductImage
	Name           string `gorm:"size:256;not null"`
	Desc           string `gorm:"size:256;not null"`
	Stock          int
	Price          decimal.Decimal `gorm:"type:decimal(16,2)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
