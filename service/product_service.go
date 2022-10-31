package service

import "mangojek-backend/model"

type ProductService interface {
	GetProduct(id int) model.GetProductResponse
	GetProducts() []model.GetProductResponse
	Save(request model.CreateProductRequest) model.CreateProductResponse
	// Delete(id int)
}
