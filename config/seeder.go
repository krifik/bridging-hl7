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
		// {Seeder: faker.ProductFaker(db)},
		{Seeder: faker.CategoryFaker(db)},
		{Seeder: faker.ProductImageFaker(db)},
		{Seeder: faker.PartnerFaker(db)},
	}
}
