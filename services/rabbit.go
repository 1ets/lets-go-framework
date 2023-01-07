package services

import (
	"lets-go-framework/app/controllers"
	"lets-go-framework/lets/drivers"
)

func RabbitEventHandler(r *drivers.MessageEngine) {
	r.Event("transfer-result", controllers.RabbitTransferResult)
}
