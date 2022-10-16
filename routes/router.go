package routes

import (
	"mangojek-backend/controller"
	"mangojek-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, controller controller.UserController) {
	app.Get("/api/users", middleware.AuthMiddleware, controller.FindAll)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Post("/api/test", controller.TestRawSQL)
	// app.Post("/api/users", middleware.AuthMiddleware, controller.Insert)
}
