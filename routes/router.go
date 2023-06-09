package routes

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/krifik/bridging-hl7/controller"
	_ "github.com/krifik/bridging-hl7/docs"
	"github.com/krifik/bridging-hl7/exception"
	"github.com/krifik/bridging-hl7/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Route(app *fiber.App, fileController controller.FileController) {
	app.Get("/api/docs/*", swagger.HandlerDefault)
	app.Get("/api/file", fileController.GetContentFile)
	app.Get("/api/files", fileController.GetFiles)
	app.Post("/api/result", fileController.CreateFileResult)
	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("HALLO")
	})
	app.Get("/api/send", func(c *fiber.Ctx) error {
		var request model.JSONRequest
		err := c.BodyParser(&request)
		exception.SendLogIfErorr(err, "29")
		jsonData, errJson := json.Marshal(request)
		exception.SendLogIfErorr(errJson, "31")
		message := amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonData,
		}
		amqpServerURL := os.Getenv("AMQP_URL")

		// Create a new RabbitMQ connection.
		connectRabbitMQ, err := amqp.Dial(amqpServerURL)
		if err != nil {
			panic(err)
		}
		defer connectRabbitMQ.Close()
		channelRabbitMQ, err := connectRabbitMQ.Channel()
		if err != nil {
			panic(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		defer channelRabbitMQ.Close()
		// Attempt to publish a message to the queue.
		if err := channelRabbitMQ.PublishWithContext(
			ctx,
			"",                // exchange
			"bridging_result", // queue name
			false,             // mandatory
			false,             // immediate
			message,           // message to publish
		); err != nil {
			return err
		}

		return nil
	})
}
