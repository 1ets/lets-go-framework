package clients

import (
	"fmt"
	"lets-go-framework/adapters/data"
	"lets-go-framework/lets/drivers"
	"os"
)

var RabbitBalance = rabbitBalance{}

type rabbitBalance struct {
	Driver *drivers.ServiceRabbit
}

func (r *rabbitBalance) BalanceTransfer(correlationId string, data *data.EventTransfer) error {
	rabbit := r.Driver

	event := drivers.Event{
		Exchange:   rabbit.Consumer.GetExchange(),
		RoutingKey: os.Getenv("RQ_ROUTING_KEY_TRANSFER"),
		ReplyTo:    rabbit.Consumer.GenerateReplyTo(),
		Body: drivers.MessageBody{
			Event: "balance-transfer",
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

func (r *rabbitBalance) BalanceRollback(data *data.EventTransfer) error {
	rabbit := r.Driver

	event := drivers.Event{
		Exchange:   rabbit.Consumer.GetExchange(),
		RoutingKey: os.Getenv("RQ_ROUTING_KEY_TRANSFER"),
		Body: drivers.MessageBody{
			Event: "balance-transfer-rollback",
			Data:  data,
		},
	}

	err := rabbit.Publish(event)
	if err != nil {
		fmt.Println("Cant publish to rabbit: ", err.Error())
	}

	return nil
}
