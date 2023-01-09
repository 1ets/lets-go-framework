package services

import (
	"lets-go-framework/adapters/servers"
	"lets-go-framework/lets/drivers"
)

func RabbitEventHandler(r *drivers.MessageEngine) {
	r.Event("transfer-result", servers.RabbitTransferResult)
	r.Event("balance-transfer-result", servers.RabbitBalanceResult)
}
