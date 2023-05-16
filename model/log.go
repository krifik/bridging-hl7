package model

import (
	"time"
)

type Log struct {
	Name        string
	Description string
	DateTime    time.Time
	Type        string
}
