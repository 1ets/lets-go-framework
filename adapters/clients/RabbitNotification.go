package clients

import (
	"fmt"
	"lets-go-framework/adapters/data"
	"lets-go-framework/lets/drivers"
	"os"
)

var RabbitNotification = rabbitNotification{}

type rabbitNotification struct {
	Driver *drivers.ServiceRabbit
}

func (r *rabbitNotification) Notify(data *data.EventNotification) error {
	rabbit := r.Driver

	event := drivers.Event{
		Exchange:   os.Getenv("RQ_EXCHANGE_NOTIFICATION"),
		RoutingKey: os.Getenv("RQ_ROUTING_KEY_NOTIFICATION"),
		Body: drivers.MessageBody{
			Event: "transfer",
			Data:  data,
		},
	}

	err := rabbit.Publish(event)
	if err != nil {
		fmt.Println("Cant publish to rabbit: ", err.Error())
	}

	return nil
}
