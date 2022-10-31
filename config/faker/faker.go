package faker

import (
	"mangojek-backend/entity"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func PartnerFaker(db *gorm.DB) *entity.Partner {
	return &entity.Partner{
		Location:  "-6.36967, 108.37207",
		Name:      "warung jajal",
		Desc:      "jajal dikit",
		OwnerName: "anonym",
		Phone:     "08123123123",
	}
}

func CategoryFaker(db *gorm.DB) *entity.Category {
	return &entity.Category{
		// ProductId: productId,
		Name: "Makanan",
	}
}
func ProductImageFaker(db *gorm.DB) *entity.ProductImage {
	return &entity.ProductImage{
		Path: "app/storage/doge.png",
		// ProductId: productId,
		Large:  "1",
		Medium: "1",
		Small:  "1",
	}
}
func ProductFaker(db *gorm.DB) *entity.Product {
	return &entity.Product{
		Name:  "Seblak",
		Desc:  "Seblack yang lebat dan berbiji",
		Stock: 69,
		Price: decimal.NewFromFloatWithExponent(69.000, 0),
	}
}
