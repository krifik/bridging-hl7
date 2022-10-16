package main

import (
	"flag"
	"mangojek-backend/app"
	"mangojek-backend/exception"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		app.InitializeDB()
		return
	}
	app := app.InitializedApp()
	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
