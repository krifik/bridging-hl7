package repository

import (
	"errors"
	"mangojek-backend/config"
	"mangojek-backend/entity"
	"mangojek-backend/exception"
	"mangojek-backend/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repository *UserRepositoryImpl) FindAll() ([]entity.User, error) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	var items []entity.User
	result := repository.DB.WithContext(ctx).Find(&items)

	if result.RowsAffected < 0 {
		return nil, errors.New("User not found")
	}
	var users []entity.User
	for _, item := range items {
		users = append(users, entity.User{
			Id:       item.Id,
			Name:     item.Name,
			Email:    item.Email,
			Password: item.Password,
		})
	}
	return users, nil
}

func (repository *UserRepositoryImpl) Delete(db *gorm.DB, userId int) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	result := db.WithContext(ctx).Delete(&userId)
	exception.PanicIfNeeded(result.Error)
}
func (repository *UserRepositoryImpl) Insert(request model.CreateUserRequest) (user entity.User) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	result := repository.DB.WithContext(ctx).Create(&entity.User{
		Id:       request.Id,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	})
	exception.PanicIfNeeded(result.Error)
	return user
}
