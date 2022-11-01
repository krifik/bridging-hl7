package service

import (
	"mangojek-backend/exception"
	"mangojek-backend/model"
	"mangojek-backend/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductServiceImpl(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepository: productRepository}
}
func (service *ProductServiceImpl) GetProduct(id int) model.GetProductResponse {
	product, err := service.ProductRepository.GetProduct(id)
	exception.PanicIfNeeded(err)
	responses := model.GetProductResponse{
		ID:             product.ID,
		Name:           product.Name,
		Desc:           product.Desc,
		Stock:          product.Stock,
		Price:          product.Price,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
		DeletedAt:      product.DeletedAt,
		ProductImageId: product.ProductImageId,
		ProductImage:   product.ProductImage,
		CategoryId:     product.CategoryId,
		Category:       product.Category,
		PartnerId:      product.PartnerId,
		Partner:        product.Partner,
	}

	return responses
}

func (service *ProductServiceImpl) GetProducts() []model.GetProductResponse {
	products, err := service.ProductRepository.GetProducts()
	exception.PanicIfNeeded(err)
	var responses []model.GetProductResponse
	for _, product := range products {
		responses = append(responses, model.GetProductResponse{
			ID:             product.ID,
			Name:           product.Name,
			Desc:           product.Desc,
			Stock:          product.Stock,
			Price:          product.Price,
			CreatedAt:      product.CreatedAt,
			UpdatedAt:      product.UpdatedAt,
			DeletedAt:      product.DeletedAt,
			ProductImageId: product.ProductImageId,
			ProductImage:   product.ProductImage,
			CategoryId:     product.CategoryId,
			Category:       product.Category,
			PartnerId:      product.PartnerId,
			Partner:        product.Partner,
		})
	}
	return responses
}

func (service *ProductServiceImpl) Save(request model.CreateProductRequest) model.CreateProductResponse {
	// validation here
	product := service.ProductRepository.Save(request)
	response := model.CreateProductResponse{
		ID:             product.ID,
		Name:           product.Name,
		Desc:           product.Desc,
		Stock:          product.Stock,
		Price:          product.Price,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
		DeletedAt:      product.DeletedAt,
		ProductImageId: product.ProductImageId,
		// ProductImage:   product.ProductImage,
		CategoryId: product.CategoryId,
		// Category:       product.Category,
		PartnerId: product.PartnerId,
		// Partner:        product.Partner,
	}
	return response
}
