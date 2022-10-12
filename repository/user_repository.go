package repository

import (
	"mangojek-backend/entity"
	"mangojek-backend/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(request model.CreateUserRequest) entity.User
	FindAll() ([]entity.User, error)
	Delete(db *gorm.DB, userId int)
}
