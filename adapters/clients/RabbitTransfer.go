package clients

import (
	"fmt"
	"lets-go-framework/adapters/data"
	"lets-go-framework/lets/rabbitmq"
	"os"
)

var RabbitTransfer = rabbitTransfer{}

type rabbitTransfer struct {
	Driver rabbitmq.RabbitClient
}

func (r *rabbitTransfer) Transfer(correlationId string, data *data.EventTransfer) error {
	rabbit := r.Driver

	event := rabbitmq.Event{
		Name:          "transfer",
		Data:          data,
		CorrelationId: correlationId,
		Exchange:      rabbit.GetDst().GetExchange(),
		RoutingKey:    os.Getenv("RQ_ROUTING_KEY_TRANSFER"),
		ReplyTo: rabbitmq.ReplyTo{
			Exchange:      os.Getenv("RQ_ROUTING_KEY_TRANSFER"),
			RoutingKey:    os.Getenv("RQ_ROUTING_KEY_TRANSFER"),
			CorrelationId: correlationId,
		},
	}

	err := rabbit.Publish(event)
	if err != nil {
		fmt.Println("Cant publish to rabbit: ", err.Error())
	}

	return nil
}

func (r *rabbitTransfer) TransferRollback(data *data.EventTransferRollback) error {
	rabbit := r.Driver

	event := rabbitmq.Event{
		Name:       "transfer-rollback",
		Data:       data,
		Exchange:   rabbit.GetDst().GetExchange(),
		RoutingKey: os.Getenv("RQ_ROUTING_KEY_TRANSFER"),
	}

	err := rabbit.Publish(event)
	if err != nil {
		fmt.Println("Cant publish to rabbit: ", err.Error())
	}

	return nil
}
