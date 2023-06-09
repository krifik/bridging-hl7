package app

import (
	"github.com/krifik/bridging-hl7/config"
	"github.com/krifik/bridging-hl7/module"
	"github.com/krifik/bridging-hl7/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitializedApp() *fiber.App {
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New(), cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// db := config.InitializedSqlite()

	// Setup Routing
	fileController := module.NewFileModule()
	routes.Route(app, fileController)
	return app

}
