package entity

import (
	"time"

	"gorm.io/gorm"
)

type Driver struct {
	ID             int    `gorm:"primaryKey,not null,autoIncrement;uniqueIndex"`
	Name           string `gorm:"size:256,not null"`
	Address        string `gorm:"size:256,not null"`
	ActivityStatus bool
	VehicleBrand   string `gorm:"size:100,not null"`
	VehicleType    string `gorm:"size:100,not null"`
	LicensePlate   string `gorm:"size:50,not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
