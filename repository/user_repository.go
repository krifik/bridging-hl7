package repository

import (
	"mangojek-backend/entity"
	"mangojek-backend/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Login(request model.CreateUserRequest) (user entity.User, err error)
	Register(request model.CreateUserRequest) (user entity.User, err error)
	FindAll() ([]entity.User, error)
	Delete(db *gorm.DB, userId int)
	CheckEmail(request model.CreateUserRequest) (result int64)
	TestRawSQL()
}
