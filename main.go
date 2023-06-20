package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
	"github.com/krifik/bridging-hl7/app"
	"github.com/krifik/bridging-hl7/config"
	_ "github.com/krifik/bridging-hl7/docs"
	"github.com/krifik/bridging-hl7/entity"
	"github.com/krifik/bridging-hl7/exception"
	"github.com/krifik/bridging-hl7/rabbitmq"
	"github.com/krifik/bridging-hl7/sftp"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		pp.Print(errEnv)
	}
	noInternet := make(chan bool)
	// var msg string
	var wg sync.WaitGroup
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	url := host + ":" + port

	// wg.Add(1)
	// go func() {
	// 	bot.StartBot()
	// 	wg.Done()
	// }()

	wg.Add(1)
	go func() {
		ch, conn := config.InitializedRabbitMQ()
		rabbitmq.StartConsumer(ch, conn)
		wg.Done()
	}()
	// wg.Add(1)
	// go func() {
	// 	watcher.StartWatcher()
	// 	wg.Done()
	// }()

	wg.Add(1)
	go func() {
		db := config.NewPostgresDatabase(config.NewConfiguration())
		entity.Migrate(db)
		sftp.Watcher(db, noInternet)
	}()

	// addr := host + ":" + port
	// Start App
	wg.Add(1)
	go func() {
		app := app.InitializedApp()
		err := app.Listen(url)
		exception.PanicIfNeeded(err)
	}()
	wg.Add(1)
	go func() {
		for {
			time.Sleep(5 * time.Second)
			_, err := net.Dial("tcp", "google.com:80")
			if err != nil {
				fmt.Println("No internet connection")
				// os.Exit(1)
				noInternet <- true
			} else {
				// conn.Close()
				noInternet <- false
				fmt.Println("Internet connection detected")
			}
		}
	}()
	wg.Wait()

}
