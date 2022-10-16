package service

import (
	"mangojek-backend/model"
)

type UserService interface {
	Insert(request model.CreateUserRequest) (response model.CreateUserResponse)
	FindAll() ([]model.GetUserResponse, error)
	Login(request model.CreateUserRequest) (response model.CreateUserResponse, err error)
	TestRawSQL()
}
