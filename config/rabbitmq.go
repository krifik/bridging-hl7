package config

import (
	"log"
	"os"

	"github.com/k0kubun/pp"
	"github.com/krifik/bridging-hl7/exception"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitializedRabbitMQ() (*amqp.Channel, *amqp.Connection) {

	err := godotenv.Load()
	exception.SendLogIfErorr(err, "12")
	url := os.Getenv("AMQP_URL")
	log.Println("Initializing RabbitMQ . . .")
	log.Println("URL : " + url)
	conn, err := amqp.Dial(url)
	if conn == nil {
		pp.Println("Reconnecting RabbitMQ . . .")
		return InitializedRabbitMQ()
	}
	exception.SendLogIfErorr(err, "17")
	ch, err := conn.Channel()
	exception.SendLogIfErorr(err, "20")
	return ch, conn
}
