package adapters

import (
	"lets-go-framework/adapters/clients"
	"lets-go-framework/lets/drivers"
)

func RabbitRegister(r *drivers.ServiceRabbit) {
	clients.RabbitTransfer.Driver = r
	clients.RabbitBalance.Driver = r
	clients.RabbitNotification.Driver = r
}
