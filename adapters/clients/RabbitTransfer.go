package clients

import (
	"fmt"
	"lets-go-framework/adapters/data"
	"lets-go-framework/lets/drivers"
	"os"
)

var RabbitTransfer = rabbitTransfer{}

type rabbitTransfer struct {
	Driver *drivers.ServiceRabbit
}

func (r *rabbitTransfer) Transfer(correlationId string, data *data.EventTransfer) error {
	rabbit := r.Driver

	event := drivers.Event{
		Exchange:   rabbit.Consumer.GetExchange(),
		RoutingKey: os.Getenv("RQ_ROUTING_KEY_TRANSFER"),
		ReplyTo:    rabbit.Consumer.GenerateReplyTo(),
		Body: drivers.MessageBody{
			Event: "transfer",
			Data:  data,
		},
		CorrelationId: correlationId,
	}

	err := rabbit.Publish(event)
	if err != nil {
		fmt.Println("Cant publish to rabbit: ", err.Error())
	}

	return nil
}

func (r *rabbitTransfer) TransferRollback(correlationId string, data *data.EventTransfer) error {
	rabbit := r.Driver

	event := drivers.Event{
		Exchange:   rabbit.Consumer.GetExchange(),
		RoutingKey: os.Getenv("RQ_ROUTING_KEY_TRANSFER"),
		ReplyTo:    rabbit.Consumer.GenerateReplyTo(),
		Body: drivers.MessageBody{
			Event: "transfer-rollback",
			Data:  data,
		},
		CorrelationId: correlationId,
	}

	err := rabbit.Publish(event)
	if err != nil {
		fmt.Println("Cant publish to rabbit: ", err.Error())
	}

	return nil
}
