package clients

import (
	"fmt"
	"lets-go-framework/adapters/data"
	"lets-go-framework/lets/rabbitmq"
	"os"
)

var RabbitBalance rabbitBalance

type rabbitBalance struct {
	Driver rabbitmq.RabbitClient
}

func (r *rabbitBalance) BalanceTransfer(correlationId string, data *data.EventTransfer) error {
	rabbit := r.Driver

	event := rabbitmq.Event{
		Name:          "balance-transfer",
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

func (r *rabbitBalance) BalanceRollback(data *data.EventTransfer) error {
	rabbit := r.Driver

	event := rabbitmq.Event{
		Name:       "balance-transfer-rollback",
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
