package config

import (
	"bridging-hl7/exception"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitializedRabbitMQ() (*amqp.Channel, *amqp.Connection) {
	log.Println("Initializing RabbitMQ . . .")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	exception.SendLogIfErorr(err, "12")

	ch, err := conn.Channel()
	exception.SendLogIfErorr(err, "15")
	return ch, conn
}
