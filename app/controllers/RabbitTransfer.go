package controllers

import (
	"lets-go-framework/app/orchestrator"
	"lets-go-framework/app/structs"

	"github.com/kataras/golog"
)

func RabbitTransferResult(correlationId string, data *structs.EventTransferResult) {
	for index, message := range orchestrator.WaitTransfer {
		if index == correlationId {
			message <- *data
			return
		}
	}
	golog.Error("Cant find correlation id")
}
