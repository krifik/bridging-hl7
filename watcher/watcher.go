package watcher

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/krifik/bridging-hl7/exception"
	"github.com/krifik/bridging-hl7/module"
	"github.com/krifik/bridging-hl7/service"
	"github.com/krifik/bridging-hl7/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/k0kubun/pp"
)

// StartWatcher starts a file system watcher and sends a JSON HTTP request
// whenever a new file is created. It uses the Fiber HTTP client to make the
// request and sends the file content as the request body. The function retrieves
// the file content using the GetContentFile method of the FileService struct.
// It also logs any errors that occur during the request.
//
// No parameters are taken in and the function does not return anything.

func StartWatcher() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	// create a new watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		pp.Print(err)
	}

	// close the watcher when the function exits
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {

					// INFO: READFILE AND CONVERT TO JSON WITH HTTP REQUEST
					client := fiber.Client{
						NoDefaultUserAgentHeader: false,
					}

					api := os.Getenv("API_EXTERNAL")

					a := client.Post(api)
					a.ContentType("application/json")

					fileService := module.UseService()
					fileContent := service.FileService.GetContentFile(fileService, event.Name)
					a.JSON(fileContent)
					var data interface{}
					code, _, errs := a.Struct(&data) // ...
					var slices []string
					for _, v := range errs {
						slices = append(slices, v.Error())
					}
					if err != nil {
						utils.SendMessage("LINE 62\n" + " Log Type: Error\n" + "Error: \n" + err.Error() + "\n")
					}
					b, err := json.MarshalIndent(fileContent, "", "    ")
					exception.SendLogIfErorr(err, "75")
					utils.SendMessage("SENT JSON" + string(b))

					errsStr := strings.Join(slices[:], "\n")
					utils.SendMessage(time.Now().Format("02-01-2006 15:04:05") + " \n\nLog Type: Info \n\n" + "API Request Message : \n\n" + "Code " + strconv.Itoa(code) + " 👍")
					if len(errs) > 0 {
						utils.SendMessage("LINE 73\n" + "Log Type: Error\n\n" + "Error: \n" + errsStr + "\n")
					}

					// NOTE: in case in the future use a rabbitmq instead REST
					// fileService := module.UseService()
					// fileContent := service.FileService.GetContentFile(fileService, event.Name)
					// err := rabbitmq.SendJsonToRabbitMQ(fileContent)
					// exception.SendLogIfErorr(err, "80")
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("File modified:", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					fmt.Println("File removed:", event.Name)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					fmt.Println("File renamed:", event.Name)
				}
				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					fmt.Println("File permissions modified:", event.Name)
				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()
	dir := os.Getenv("ORDERDIR")
	log.Println(dir)
	errFilePath := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			watcher.Add(path)
		}
		return nil
	})
	exception.SendLogIfErorr(errFilePath, "110")
	<-done
}
