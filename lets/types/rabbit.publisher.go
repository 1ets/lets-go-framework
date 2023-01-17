package types

import "github.com/kataras/golog"

// Default configuration
const (
	RQ_PUBLISHER_EXCHANGE    = "default-exchange"
	RQ_PUBLISHER_ROUTING_KEY = "default-routing-key"
	RQ_PUBLISHER_QUEUE       = "default-queue"
)

// Interface for dsn accessable method
type IRabbitPublisher interface {
	GetExchange() string
	GetRoutingKey() string
	GetQueue() string
}

// Target host information.
type RabbitPublisher struct {
	Exchange, RoutingKey, Queue string
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

// func NewRabbitPublisher(exchange, routingKey, queue string) IRabbitPublisher {
// 	return &RabbitPublisher{
// 		Exchange:   exchange,
// 		RoutingKey: routingKey,
// 		Queue:      queue,
// 	}
// }
