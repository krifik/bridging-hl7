package app

import (
	"bridging-hl7/config"
	"bridging-hl7/module"
	"bridging-hl7/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func InitializedApp() *fiber.App {
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New(), cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	db := config.InitializedSqlite()

	// Setup Routing
	fileController := module.NewFileModule(db)
	routes.Route(app, fileController)
	return app

}
