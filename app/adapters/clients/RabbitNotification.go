package clients

import (
	"fmt"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/lets/rabbitmq"
	"os"
)

var RabbitNotification = rabbitNotification{}

type rabbitNotification struct {
	Driver rabbitmq.RabbitClient
}

func (r *rabbitNotification) Notify(data *data.EventNotification) error {
	rabbit := r.Driver

	event := rabbitmq.Event{
		Name:       "transfer",
		Data:       data,
		Exchange:   rabbit.GetDst().GetExchange(),
		RoutingKey: os.Getenv("RQ_ROUTING_KEY_TRANSFER"),
		ReplyTo: rabbitmq.ReplyTo{
			Exchange:   os.Getenv("RQ_EXCHANGE_TRANSFER"),
			RoutingKey: os.Getenv("RQ_ROUTING_KEY_TRANSFER"),
		},
	}

	err := rabbit.Publish(event)
	if err != nil {
		fmt.Println("Cant publish to rabbit: ", err.Error())
	}

	return nil
}
