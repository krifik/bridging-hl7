package repository

import (
	"errors"
	"fmt"
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
		return nil, errors.New("user not found")
	}
	var users []entity.User
	for _, item := range items {
		users = append(users, entity.User{
			Id:        item.Id,
			FirstName: item.FirstName,
			LastName:  item.LastName,
			Email:     item.Email,
			Password:  item.Password,
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
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	})
	exception.PanicIfNeeded(result.Error)
	return user
}
func (repository *UserRepositoryImpl) TestRawSQL() {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	payload := struct {
		Name     string
		Email    string
		Password string
	}{
		Name:     "fikri",
		Email:    "fikri@gmail.com",
		Password: "password",
	}
	sql := fmt.Sprintf("INSERT INTO users(name,email,password) VALUES('%s', '%s', '%s')", payload.Name, payload.Email, payload.Password)
	fmt.Println(sql)
	repository.DB.WithContext(ctx).Create(&entity.User{
		Name:     "fikri",
		Email:    "fikri@gmail.com",
		Password: "password",
	})
	// exception.PanicIfNeeded(err)
	// id, err := result.LastInsertId()
	// exception.PanicIfNeeded(err)
	// user := entity.User{
	// 	ID
	// }
	// return result
}
