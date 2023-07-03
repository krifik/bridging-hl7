package entity

import (
	"time"

	"github.com/k0kubun/pp"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	FileName     string `gorm:"unique;size:256"`
	ReadState    bool   `gorm:"default:false"`
	ModifiedTime time.Time
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&File{})
	if err != nil {
		panic(err)
	}
	pp.Println("Migrate")
}
