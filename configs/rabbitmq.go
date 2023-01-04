package configs

import (
	"lets-go-framework/lets/types"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

// Initialize global variables
var RabbitDsn types.IRabbitDsn
var RabbitConsumer types.IRabbitConsumer

// Callable function for initialize all libraries configuration.
func Initialize() {
	RabbitDsn = &types.RabbitDsn{
		Host:        os.Getenv("RQ_HOST"),
		Port:        os.Getenv("RQ_PORT"),
		Username:    os.Getenv("RQ_USERNAME"),
		Password:    os.Getenv("RQ_PASSWORD"),
		VirtualHost: os.Getenv("RQ_VHOST"),
	}

	RabbitConsumer = &types.RabbitConsumer{
		Name:         os.Getenv("RQ_NAME"),
		Exchange:     os.Getenv("RQ_EXCHANGE"),
		ExchangeType: amqp091.ExchangeDirect,
		RoutingKey:   os.Getenv("RQ_ROUTING_KEY"),
		Queue:        os.Getenv("RQ_QUEUE"),
	}
}
