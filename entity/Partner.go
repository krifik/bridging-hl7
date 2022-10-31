package entity

import (
	"time"

	"gorm.io/gorm"
)

type Partner struct {
	ID        int `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	ProductId int `gorm:"index"`
	// Product   []Product     `gorm:"foreignKey:ID;references:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Location  string `gorm:"size:100"`
	Name      string `gorm:"size:100"`
	Desc      string `gorm:"size:256"`
	OwnerName string `gorm:"size:100"`
	Phone     string `gorm:"size:30"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
