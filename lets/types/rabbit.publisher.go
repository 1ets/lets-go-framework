package types

import (
	"encoding/json"
	"fmt"

	"github.com/kataras/golog"
)

// Default configuration
const (
	RQ_PUBLISHER_NAME        = "default-name"
	RQ_PUBLISHER_EXCHANGE    = "default-exchange"
	RQ_PUBLISHER_ROUTING_KEY = "default-routing-key"
	RQ_PUBLISHER_QUEUE       = "default-queue"
)

// Interface for dsn accessable method
type IRabbitMQPublisher interface {
	GetName() string
	GetExchange() string
	GetRoutingKey() string
	GetQueue() string
	GetReplyTo() ReplyTo
}

// Target host information.
type RabbitPublisher struct {
	Name, Exchange, RoutingKey, Queue string
	ReplyTo                           ReplyTo
}

// Get ExchangeName.
func (rtm *RabbitPublisher) GetName() string {
	if rtm.Name == "" {
		golog.Warn("Configs RabbitMQ: RQ_PUBLISHER_NAME is not set in .env file, using default configuration.")

		return RQ_PUBLISHER_NAME
	}

	return rtm.Name
}

// Get ExchangeName.
func (rtm *RabbitPublisher) GetExchange() string {
	if rtm.Exchange == "" {
		golog.Warn("Configs RabbitMQ: RQ_PUBLISHER_EXCHANGE is not set in .env file, using default configuration.")

		return RQ_PUBLISHER_EXCHANGE
	}

	return rtm.Exchange
}

// Get QueueName.
func (rtm *RabbitPublisher) GetRoutingKey() string {
	if rtm.RoutingKey == "" {
		golog.Warn("Configs RabbitMQ: RQ_PUBLISHER_ROUTING_KEY is not set in .env file, using default configuration.")

		return RQ_PUBLISHER_ROUTING_KEY
	}

	return rtm.RoutingKey
}

func (rtm *RabbitPublisher) GetQueue() string {
	if rtm.Queue == "" {
		golog.Warn("Configs RabbitMQ: RQ_PUBLISHER_QUEUE is not set in .env file, using default configuration.")

		return RQ_PUBLISHER_QUEUE
	}

	return rtm.Queue
}

func (rtm *RabbitPublisher) GenerateReplyTo() string {
	replyTo := map[string]string{
		"exchange":    rtm.GetExchange(),
		"routing_key": rtm.GetRoutingKey(),
	}

	_replyTo, err := json.Marshal(replyTo)
	if err != nil {
		fmt.Println("cant marshal replyTos")
	}

	return string(_replyTo)
}
