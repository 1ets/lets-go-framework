package controllers

import (
	"fmt"
	"lets-go-framework/app/orchestrator"
	"lets-go-framework/app/structs"

	"github.com/kataras/golog"
)

func RabbitTransferResult(correlationId string, data *structs.EventTransferResult) {
	golog.Info("r.GetCorrelationId() ", correlationId)

	for index, message := range orchestrator.WaitTransferStart {
		golog.Infof("index: %v", index)
		golog.Infof("message: %v", message)

		if index == correlationId {
			fmt.Println("PREP PUSH")
			// golog.Info(string(r.GetData()))
			// var message structs.EventTransferResult
			// err := json.Unmarshal(data, &message)
			// if err != nil {
			// 	golog.Errorf("broken: ", err.Error())
			// }
			fmt.Println("PUSH")
			message <- *data

			fmt.Println("END PUSH")
		}
	}
}
