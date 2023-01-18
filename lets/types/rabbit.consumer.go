package types

import (
	"encoding/json"
	"fmt"

	"github.com/kataras/golog"
)

// Default configuration
const (
	RQ_CONSUMER_EXCHANGE    = "default-exchange"
	RQ_CONSUMER_ROUTING_KEY = "default-routing-key"
	RQ_CONSUMER_QUEUE       = "default-queue"
)

// Interface for dsn accessable method
type IRabbitMQConsumer interface {
	GetName() string
	GetExchange() string
	GetExchangeType() string
	GetRoutingKey() string
	GetQueue() string
}

// Target host information.
type RabbitMQConsumer struct {
	Name, Exchange, ExchangeType, RoutingKey, Queue string
}

// Get ExchangeName.
func (rtm *RabbitMQConsumer) GetName() string {
	return rtm.Name
}

// Get ExchangeName.
func (rtm *RabbitMQConsumer) GetExchange() string {
	if rtm.Exchange == "" {
		golog.Warn("Configs RabbitMQ: RQ_CONSUMER_EXCHANGE is not set in .env file, using default configuration.")

		return RQ_CONSUMER_EXCHANGE
	}

	return rtm.Exchange
}

// Get ExchangeName.
func (rtm *RabbitMQConsumer) GetExchangeType() string {
	return rtm.ExchangeType
}

// Get QueueName.
func (rtm *RabbitMQConsumer) GetRoutingKey() string {
	if rtm.RoutingKey == "" {
		golog.Warn("Configs RabbitMQ: RQ_CONSUMER_ROUTING_KEY is not set in .env file, using default configuration.")

		return RQ_CONSUMER_ROUTING_KEY
	}

	return rtm.RoutingKey
}

func (rtm *RabbitMQConsumer) GetQueue() string {
	if rtm.Queue == "" {
		golog.Warn("Configs RabbitMQ: RQ_CONSUMER_QUEUE is not set in .env file, using default configuration.")

		return RQ_CONSUMER_QUEUE
	}

	return rtm.Queue
}

func (rtm *RabbitMQConsumer) GenerateReplyTo() string {
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
