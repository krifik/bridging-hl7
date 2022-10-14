package service

import (
	"mangojek-backend/model"
)

type UserService interface {
	Insert(request model.CreateUserRequest) (response model.CreateUserResponse)
	FindAll() ([]model.GetUserResponse, error)
	// Delete(db *gorm.DB, userId int)
}
