package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID      int     `gorm:"primaryKey,not null,autoIncrement;uniqueIndex;"`
	OrderId int     `gorm:"unique"`
	Order   []Order `gorm:"foreignKey:ID;references:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name    string  `gorm:"size:256"`
	// LastName  string  `gorm:"size:256"`
	Email    string `gorm:"size:256"`
	Password string `gorm:"size:256"`
	// RememberToken string  `gorm:"size:256"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
