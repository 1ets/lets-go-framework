package controllers

import (
	"lets-go-framework/app/orchestrator"
	"lets-go-framework/app/structs"

	"github.com/kataras/golog"
)

func RabbitBalanceResult(correlationId string, data *structs.EventBalanceTransferResult) {
	for index, message := range orchestrator.WaitBalanceTransfer {
		if index == correlationId {
			message <- *data
			return
		}
	}
	golog.Error("Cant find correlation id")
}

func RabbitBalanceRollbackResult(correlationId string, data *structs.EventTransferResult) {
	for index, message := range orchestrator.WaitTransfer {
		if index == correlationId {
			message <- *data
			return
		}
	}
	golog.Error("Cant find correlation id")
}
