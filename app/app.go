package app

import (
	"mangojek-backend/config"
	"mangojek-backend/controller"
	"mangojek-backend/repository"
	"mangojek-backend/routes"
	"mangojek-backend/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitializedApp() *fiber.App {
	configuration := config.NewConfiguration()
	database := config.NewPostgresDatabase(configuration)
	config.NewRunMigration(database)

	// Setup Repository
	userRepository := repository.NewUserRepositoryImpl(database)

	// Setup Service
	userService := service.NewUserServiceImpl(userRepository)

	// Setup Controller
	userController := controller.NewUserControllerImpl(userService)

	// Setup Fiber
	// app := fiber.New(config.NewFiberConfig())
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	routes.Route(app, userController)
	return app

}
