package entity

import (
	"time"

	"gorm.io/gorm"
)

type ProductImage struct {
	Id        int    `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	ProductId int    `gorm:"index"`
	Path      string `gorm:"type:text"`
	Large     string `gorm:"type:text"`
	Medium    string `gorm:"type:text"`
	Small     string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
