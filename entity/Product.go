package entity

import (
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID             int           `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	CategoryId     sql.NullInt64 `gorm:"index"`
	Category       Category      `gorm:"foreignKey:CategoryId;references:ID;"`
	PartnerId      sql.NullInt64 `gorm:"index"`
	Partner        Partner       `gorm:"foreignKey:PartnerId;references:ID;"`
	ProductImageId sql.NullInt64 `gorm:"index"`
	ProductImage   ProductImage  `gorm:"foreignKey:ProductImageId;references:ID;"`
	Name           string        `gorm:"size:256;not null"`
	Desc           string        `gorm:"size:256;not null"`
	Stock          int
	Price          decimal.Decimal `gorm:"type:decimal(16,2)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
