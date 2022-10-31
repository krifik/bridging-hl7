package repository

import (
	"mangojek-backend/entity"
	"mangojek-backend/model"
)

type ProductRepository interface {
	GetProduct(id int) (entity.Product, error)
	GetProducts() ([]entity.Product, error)
	Save(request model.CreateProductRequest) entity.Product
	Delete(id int)
}
