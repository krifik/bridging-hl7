package entity

import (
	"time"
)

type File struct {
	ID        int    `gorm:"primaryKey,not null,autoIncrement;uniqueIndex;"`
	FileName  string `gorm:"size:256"`
	CreatedAt time.Time
}
