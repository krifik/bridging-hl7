package repository

import (
	"database/sql"
	"mangojek-backend/config"
	"mangojek-backend/entity"
	"mangojek-backend/exception"
	"mangojek-backend/model"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}
func (repository *ProductRepositoryImpl) GetProduct(id int) (entity.Product, error) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	var product entity.Product
	err := repository.DB.WithContext(ctx).First(&product).Error
	exception.PanicIfNeeded(err)
	return product, nil
}

func (repository *ProductRepositoryImpl) GetProducts() ([]entity.Product, error) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	var products []entity.Product
	err := repository.DB.WithContext(ctx).Find(&products).Error
	exception.PanicIfNeeded(err)
	return products, nil
}

func (repository *ProductRepositoryImpl) Save(request model.CreateProductRequest) entity.Product {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	product := entity.Product{
		ID: request.ID,
		ProductImageId: sql.NullInt64{
			Int64: int64(request.ProductImageId),
		},
		// ProductImage:   request.ProductImage,
		CategoryId: sql.NullInt64{
			Int64: int64(request.CategoryId),
		},
		// Category:       request.Category,
		PartnerId: sql.NullInt64{
			Int64: int64(request.PartnerId),
		},
		// Partner:   request.Partner,
		Name:      request.Name,
		Desc:      request.Desc,
		Stock:     request.Stock,
		Price:     request.Price,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
		DeletedAt: request.DeletedAt,
	}

	repository.DB.WithContext(ctx).Omit("Partner", "Category", "ProductImage").Create(&product)
	return product
}

func (repository *ProductRepositoryImpl) Delete(id int) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	repository.DB.WithContext(ctx).Delete(id)
}
