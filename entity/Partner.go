package entity

import (
	"time"

	"gorm.io/gorm"
)

type Partner struct {
	Id        int       `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	ProductId int       `gorm:"unique"`
	Product   []Product `gorm:"foreignKey:Id;references:ProductId"`
	Location  string    `gorm:"size:100"`
	Name      string    `gorm:"size:100"`
	Desc      string    `gorm:"size:256"`
	OwnerName string    `gorm:"size:100"`
	Phone     string    `gorm:"size:30"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
