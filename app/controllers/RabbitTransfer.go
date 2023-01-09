package controllers

import (
	"lets-go-framework/app/orchestrator"
	"lets-go-framework/app/structs"
)

func RabbitTransferResult(correlationId string, data *structs.EventTransferResult) {
	for index, message := range orchestrator.WaitTransferStart {
		if index == correlationId {
			message <- *data
		}
	}
}
