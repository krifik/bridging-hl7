package main

import (
	"bridging-hl7/app"
	"bridging-hl7/bot"
	"bridging-hl7/config"
	_ "bridging-hl7/docs"
	"bridging-hl7/exception"
	"bridging-hl7/model"
	"bridging-hl7/rabbitmq"
	"bridging-hl7/watcher"
	"os"
	"sync"

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
		watcher.StartWatcher()
		wg.Done()
	}()

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
