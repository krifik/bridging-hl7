package main

import (
	"mangojek-backend/app"
	"mangojek-backend/exception"
)

func main() {
	app := app.InitializedApp()
	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
