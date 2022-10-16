package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id        int       `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	Name      string    `gorm:"size:100"`
	ProductId int       `gorm:"unique"`
	Product   []Product `gorm:"foreignKey:Id;references:ProductId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
