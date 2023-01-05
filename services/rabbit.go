package services

import (
	"lets-go-framework/app/controllers"
	"lets-go-framework/lets"
)

func RabbitEventHandler(r *lets.MessageEngine) {
	r.Event("transfer-result", controllers.RabbitTransferResult)
}
