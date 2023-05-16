package rabbitmq

import (
	"bridging-hl7/exception"
	"bridging-hl7/model"
	"bridging-hl7/module"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// global channel
var globalConsumer chan struct{}

func StartConsumer(ch *amqp.Channel, conn *amqp.Connection) {
	defer ch.Close()
	defer conn.Close()
	log.Println("Starting Consumer")
	q, err := ch.QueueDeclare(
		"coba", // name
		true,   // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	exception.SendLogIfErorr(err, "20")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	exception.SendLogIfErorr(err, "30")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	exception.SendLogIfErorr(err, "37")
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	// Make a channel to receive messages into infinite loop.
	forever := make(chan struct{})
	fileService := module.UseService()
	go func() {
		for message := range msgs {
			var request model.JSONRequest
			// For example, show received message in a console.
			err := json.Unmarshal(message.Body, &request)
			exception.SendLogIfErorr(err, "57")
			fileService.CreateFileResult(request)
			log.Printf(" > Received message: %s\n", message.Body)
		}
	}()
	<-forever
}
