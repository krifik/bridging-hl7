package main

import (
	"flag"
	"mangojek-backend/app"
	"mangojek-backend/exception"
)

func main() {
	flag.Parse()
	if arg := flag.Arg(0); arg != "" {
		app.InitializeDB()
		return
	}
	app := app.InitializedApp()
	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
