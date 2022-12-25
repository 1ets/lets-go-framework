package controllers

import (
	"lets-go-framework/configs"
	"lets-go-framework/helpers"
)

func CreateOrder() {
	credential := configs.RabbitMQCredentials["order"]

	rm := helpers.RmqPublish{}
	rm.SetCredential(credential).SetEvent("request.create.order")
	rm.SetData("DATA-DATA")
	rm.Publish(false)
}
