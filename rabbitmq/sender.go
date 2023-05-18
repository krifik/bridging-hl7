package rabbitmq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/krifik/bridging-hl7/config"
	"github.com/krifik/bridging-hl7/exception"
	"github.com/krifik/bridging-hl7/model"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SendJsonToRabbitMQ(request model.Json) error {

	ch, conn := config.InitializedRabbitMQ()
	defer ch.Close()
	defer conn.Close()

	jsonData, errJson := json.Marshal(request)
	exception.SendLogIfErorr(errJson, "22")
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonData,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer ch.Close()
	// Attempt to publish a message to the queue.
	if err := ch.PublishWithContext(
		ctx,
		"",               // exchange
		"bridging_order", // queue name
		false,            // mandatory
		false,            // immediate
		message,          // message to publish
	); err != nil {
		return err
	}

	return nil
}
