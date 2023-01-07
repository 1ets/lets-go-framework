package adapters

import "github.com/rabbitmq/amqp091-go"

func RabbitRegister(c *amqp091.Connection) {
	eventTransaction.Connection = c
}
