package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	OrderId int `gorm:"unique"`
	// Order   []Order `gorm:"foreignKey:Id;references:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name string `gorm:"size:256"`
	// LastName  string  `gorm:"size:256"`
	Email    string `gorm:"size:256"`
	Password string `gorm:"size:256"`
	// RememberToken string  `gorm:"size:256"`
}
