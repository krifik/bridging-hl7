package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        int       `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	Name      string    `gorm:"size:100"`
	ProductId int       `gorm:"unique"`
	Product   []Product `gorm:"foreignKey:ID;references:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
