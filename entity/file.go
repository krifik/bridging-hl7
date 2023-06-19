package entity

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	FileName  string `gorm:"unique;size:256"`
	ReadState bool   `gorm:"default:false"`
}
