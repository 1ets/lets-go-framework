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
