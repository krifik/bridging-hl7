package entity

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	ID        int           `gorm:"primaryKey,not null,autoIncrement;uniqueIndex;"`
	ProductId sql.NullInt64 `gorm:"index"`
	Product   []Product     `gorm:"foreignKey:ID;references:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Path      string        `gorm:"type:text"`
	Large     string        `gorm:"type:text"`
	Medium    string        `gorm:"type:text"`
	Small     string        `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
