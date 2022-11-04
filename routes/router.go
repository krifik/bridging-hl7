package routes

import (
	"mangojek-backend/controller"
	_ "mangojek-backend/docs"
	"mangojek-backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Route(app *fiber.App, userController controller.UserController, productController controller.ProductController) {
	app.Get("/api/users", middleware.AuthMiddleware, userController.FindAll)
	app.Post("/api/register", userController.Register)
	app.Post("/api/login", userController.Login)
	app.Post("/api/test", userController.TestRawSQL)
	app.Get("/api/product/:id", productController.GetProduct)
	app.Post("/api/product", productController.Save)
	app.Get("/api/products", productController.GetProducts)
	app.Get("/api/docs/*", swagger.HandlerDefault)
	// app.Post("/api/users", middleware.AuthMiddleware, controller.Insert)
}
