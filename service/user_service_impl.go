package service

import (
	"mangojek-backend/entity"
	"mangojek-backend/model"
	"mangojek-backend/repository"
	"mangojek-backend/validation"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (service *UserServiceImpl) FindAll() ([]model.GetUserResponse, error) {
	users, _ := service.UserRepository.FindAll()
	var responses []model.GetUserResponse
	for _, user := range users {
		responses = append(responses, model.GetUserResponse{
			Id:       user.Id,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
	}
	return responses, nil
}

func (service *UserServiceImpl) Insert(request model.CreateUserRequest) (response model.CreateUserResponse) {
	validation.Validate(request)
	user := entity.User{
		Id:       request.Id,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	service.UserRepository.Insert(request)
	response = model.CreateUserResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	return response
}
