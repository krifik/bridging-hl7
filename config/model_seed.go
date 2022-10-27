package config

import (
	"mangojek-backend/config/faker"

	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func RegisterSeeder(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: faker.PartnerFaker(db)},
	}
}
