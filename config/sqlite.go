package config

import (
	"bridging-hl7/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializedSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("files.db"), &gorm.Config{})
	if err != nil {
		utils.SendMessage("SQLite ERROR: " + err.Error())
	}
	return db
}
