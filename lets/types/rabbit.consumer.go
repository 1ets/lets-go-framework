package types

import (
	"lets-go-framework/lets"

	"github.com/rabbitmq/amqp091-go"
)

// Default configuration
const (
	LISTEN_RABBIT_NAME          = "Default Manager"
	LISTEN_RABBIT_VHOST         = "/"
	LISTEN_RABBIT_EXCHANGE      = "lets"
	LISTEN_RABBIT_EXCHANGE_TYPE = amqp091.ExchangeDirect
	LISTEN_RABBIT_ROUTING_KEY   = "callback"
	LISTEN_RABBIT_QUEUE         = "lets.callback"
	LISTEN_RABBIT_DEBUG         = true
)

// Interface for dsn accessable method
type IRabbitMQConsumer interface {
	GetName() string
	GetExchange() string
	GetExchangeType() string
	GetRoutingKey() string
	GetQueue() string
	GetDebug() bool
	GetListener() func(Engine)
	GenerateReplyTo() ReplyTo
}

// Target host information.
type RabbitMQConsumer struct {
	Name         string `json:"name"`
	Exchange     string `json:"exchange"`
	ExchangeType string `json:"type"`
	RoutingKey   string `json:"routing_key"`
	Queue        string `json:"queue"`
	Debug        string `json:"debug"`
	Listener     func(Engine)
}

// Get Name.
func (r *RabbitMQConsumer) GetName() string {
	if r.Name == "" {
		lets.LogW("Configs: LISTEN_RABBIT_NAME is not set, using default configuration.")

		return LISTEN_RABBIT_NAME
	}
	return r.Name
}

// Get Exchange.
func (r *RabbitMQConsumer) GetExchange() string {
	if r.Exchange == "" {
		lets.LogW("Configs: LISTEN_RABBIT_EXCHANGE is not set, using default configuration.")

		return LISTEN_RABBIT_EXCHANGE
	}

	return r.Exchange
}

// Get Exchange Type.
func (r *RabbitMQConsumer) GetExchangeType() string {
	if r.ExchangeType == "" {
		lets.LogW("Config: LISTEN_RABBIT_EXCHANGE_TYPE is not set, using default configuration.")

		return LISTEN_RABBIT_EXCHANGE_TYPE
	}

	return r.ExchangeType
}

// Get Routing Key.
func (r *RabbitMQConsumer) GetRoutingKey() string {
	if r.RoutingKey == "" {
		lets.LogW("Configs: LISTEN_RABBIT_ROUTING_KEY is not set, using default configuration.")

		return LISTEN_RABBIT_ROUTING_KEY
	}

	return r.RoutingKey
}

// Get Queue.
func (r *RabbitMQConsumer) GetQueue() string {
	if r.Queue == "" {
		lets.LogW("Configs: LISTEN_RABBIT_QUEUE is not set, using default configuration.")

		return LISTEN_RABBIT_QUEUE
	}

	return r.Queue
}

// Get Debug.
func (r *RabbitMQConsumer) GetDebug() bool {
	if r.Debug == "" {
		lets.LogW("Configs: LISTEN_RABBIT_QUEUE is not set, using default configuration.")

		return LISTEN_RABBIT_DEBUG
	}

	return r.Debug == "true"
}

// Get Listener.
func (r *RabbitMQConsumer) GetListener() func(Engine) {
	return r.Listener
}

// Generating reply to payload.
func (r *RabbitMQConsumer) GenerateReplyTo() (replyTo ReplyTo) {
	lets.Bind(r, &replyTo)
	return
}
