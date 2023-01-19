package services

import (
	"lets-go-framework/app/adapters/servers"
	"lets-go-framework/lets/rabbitmq"
)

func RabbitEventHandler(r *rabbitmq.Engine) {
	r.Event("transfer-result", servers.RabbitTransferResult)
	r.Event("balance-transfer-result", servers.RabbitBalanceResult)
}
