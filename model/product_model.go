package model

import (
	"mangojek-backend/entity"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type CreateProductRequest struct {
	ID             int             `json:"id"`
	PartnerId      int             `json:"partner_id"`
	CategoryId     int             `json:"category_id"`
	ProductImageId int             `json:"product_image_id"`
	Name           string          `json:"name"`
	Desc           string          `json:"desc"`
	Stock          int             `json:"stock"`
	Price          decimal.Decimal `json:"price"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      gorm.DeletedAt  `json:"deleted_at"`
	// Partner        entity.Partner      `json:"partner"`
	// Category       entity.Category     `json:"category"`
	// ProductImage   entity.ProductImage `json:"product_image"`
}
type CreateProductResponse struct {
	ID             int             `json:"id"`
	PartnerId      int             `json:"partner_id"`
	CategoryId     int             `json:"category_id"`
	ProductImageId int             `json:"product_image_id"`
	Name           string          `json:"name"`
	Desc           string          `json:"desc"`
	Stock          int             `json:"stock"`
	Price          decimal.Decimal `json:"price"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      gorm.DeletedAt  `json:"deleted_at"`
	// Partner        entity.Partner      `json:"partner"`
	// Category       entity.Category     `json:"category"`
	// ProductImage   entity.ProductImage `json:"product_image"`
}
type GetProductResponse struct {
	ID             int                 `json:"id"`
	Name           string              `json:"name"`
	Desc           string              `json:"desc"`
	Stock          int                 `json:"stock"`
	Price          decimal.Decimal     `json:"price"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
	DeletedAt      gorm.DeletedAt      `json:"deleted_at"`
	ProductImageId int                 `json:"product_image_id"`
	ProductImage   entity.ProductImage `json:"product_image"`
	CategoryId     int                 `json:"category_id"`
	Category       entity.Category     `json:"category"`
	PartnerId      int                 `json:"partner_id"`
	Partner        entity.Partner      `json:"partner"`
}
