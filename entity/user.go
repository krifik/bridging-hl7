package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:256"`
	Email    string `gorm:"size:256"`
	Password string `gorm:"size:256"`
}
