package module

import (
	"mangojek-backend/controller"
	"mangojek-backend/repository"
	"mangojek-backend/service"

	"gorm.io/gorm"
)

func NewUserModule(database *gorm.DB) controller.UserController {
	// Setup Repository
	userRepository := repository.NewUserRepositoryImpl(database)

	// Setup Service
	userService := service.NewUserServiceImpl(userRepository)

	// Setup Controller
	userController := controller.NewUserControllerImpl(userService)
	return userController

}
func NewProductModule(database *gorm.DB) controller.ProductController {
	// Setup Repository
	productRepository := repository.NewProductRepositoryImpl(database)

	// Setup Service
	productService := service.NewProductServiceImpl(productRepository)

	// Setup Controller
	productController := controller.NewProductControllerImpl(productService)
	return productController

}
