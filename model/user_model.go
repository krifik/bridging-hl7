package model

type CreateUserRequest struct {
	Id int `json:"id"`
	// FirstName string `json:"first_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Id int `json:"id"`
	// FirstName string `json:"first_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserResponse struct {
	Id int `json:"id"`
	// FirstName string `json:"first_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
