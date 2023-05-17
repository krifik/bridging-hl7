package config

import (
	"bridging-hl7/exception"
	"log"
	"os"

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
	exception.PanicIfNeeded(err)
	exception.SendLogIfErorr(err, "17")
	ch, err := conn.Channel()
	exception.SendLogIfErorr(err, "20")
	return ch, conn
}
