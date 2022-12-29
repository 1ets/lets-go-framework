package controllers

import (
	"lets-go-framework/configs"
	"lets-go-framework/helpers"

	"github.com/gin-gonic/gin"
)

func PostTransferMoney(g *gin.Context) {
	credential := configs.RabbitMQCredentials["order"]

	rm := helpers.RmqPublish{}
	rm.SetCredential(credential).SetEvent("request.create.order")
	rm.SetData("DATA-DATA")
	rm.Publish(false)
}
func PostTransferMoneyFailed(g *gin.Context) {
	credential := configs.RabbitMQCredentials["order"]

	rm := helpers.RmqPublish{}
	rm.SetCredential(credential).SetEvent("request.create.order")
	rm.SetData("DATA-DATA")
	rm.Publish(false)
}
