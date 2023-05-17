package main

import (
	"os"
	"sync"

	"girhub.com/krifik/bridging-hl7/app"
	"girhub.com/krifik/bridging-hl7/bot"
	"girhub.com/krifik/bridging-hl7/config"
	_ "girhub.com/krifik/bridging-hl7/docs"
	"girhub.com/krifik/bridging-hl7/exception"
	"girhub.com/krifik/bridging-hl7/model"
	"girhub.com/krifik/bridging-hl7/rabbitmq"
	"girhub.com/krifik/bridging-hl7/watcher"

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
