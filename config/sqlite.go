package config

import (
	"github.com/krifik/bridging-hl7/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitializedSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("files.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		utils.SendMessage("SQLite ERROR: " + err.Error())
	}
	return db
}
