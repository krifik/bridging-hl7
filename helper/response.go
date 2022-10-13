package helper

import (
	"mangojek-backend/entity"
	"mangojek-backend/model"
)

func ToUserResponse(user entity.User) model.CreateUserRequest {
	return model.CreateUserRequest{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
func ToUserResponses(users []entity.User) []model.CreateUserRequest {
	var usersResponse []model.CreateUserRequest
	for _, user := range users {
		usersResponse = append(usersResponse, ToUserResponse(user))
	}
	return usersResponse
}
