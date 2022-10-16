package entity

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	Id             int          `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	ProductImageId int          `gorm:"unique"`
	ProductImage   ProductImage `gorm:"foreignKey:Id;references:ProductImageId"`
	CategoryId     int          `gorm:"unique"`
	Category       Category     `gorm:"foreignKey:CategoryId;references:Id"`
	PartnerId      int          `gorm:"unique"`
	Partner        Partner      `gorm:"foreignKey:PartnerId;references:Id"`
	Name           string       `gorm:"size:256;not null"`
	Desc           string       `gorm:"size:256;not null"`
	Stock          int
	Price          decimal.Decimal `gorm:"type:decimal(16,2)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
