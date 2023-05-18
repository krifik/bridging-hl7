package main

import (
	"os"
	"sync"

	"github.com/krifik/bridging-hl7/app"
	"github.com/krifik/bridging-hl7/bot"
	"github.com/krifik/bridging-hl7/config"
	_ "github.com/krifik/bridging-hl7/docs"
	"github.com/krifik/bridging-hl7/exception"
	"github.com/krifik/bridging-hl7/model"
	"github.com/krifik/bridging-hl7/rabbitmq"
	"github.com/krifik/bridging-hl7/watcher"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
)

var globalConsumer chan model.JSONRequest

func main() {

	errEnv := godotenv.Load()
	var wg sync.WaitGroup
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	url := host + ":" + port

	wg.Add(1)
	go func() {
		bot.StartBot()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		ch, conn := config.InitializedRabbitMQ()
		rabbitmq.StartConsumer(ch, conn)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		watcher.StartWatcher()
		wg.Done()
	}()
	if errEnv != nil {
		pp.Print(errEnv)
	}

	// addr := host + ":" + port
	app := app.InitializedApp()
	// Start App
	go func() {
		err := app.Listen(url)
		exception.PanicIfNeeded(err)
	}()
	wg.Wait()

}
