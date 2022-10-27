package faker

import (
	"mangojek-backend/entity"

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
