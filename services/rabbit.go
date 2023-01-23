package services

import (
	"lets-go-framework/app/adapters/clients"
	"lets-go-framework/app/adapters/servers"
	"lets-go-framework/lets/types"
)

func RabbitMQRouter(route types.Engine) {
	route.Event("example-event", servers.RabbitExample)
	route.Event("callback", clients.RabbitExample.GreetingSyncCallback)
}

func RabbitMQRouterSaga(route types.Engine) {
	route.Event("transfer-request", servers.RabbitSagaExampleTransfer)
	route.Event("balance-request", servers.RabbitSagaExampleBalance)
	route.Event("notification-request", servers.RabbitSagaExampleNotification)

	route.Event("transfer-callback", clients.RabbitSagaExample.TransferCallback)
	route.Event("balance-callback", clients.RabbitSagaExample.BalanceCallback)
	route.Event("notification-callback", clients.RabbitSagaExample.NotificationCallback)
}
