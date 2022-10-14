package config

import "mangojek-backend/entity"

type Entity struct {
	Entity interface{}
}

func RegisterEntities() []Entity {
	return []Entity{
		{Entity: entity.User{}},
		{Entity: entity.Driver{}},
		{Entity: entity.Payment{}},
		{Entity: entity.Partner{}},
		{Entity: entity.Product{}},
		{Entity: entity.ProductImage{}},
		{Entity: entity.Category{}},
		{Entity: entity.Order{}},
	}
}
