package servers

import (
	"fmt"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/controllers"
	"lets-go-framework/app/structs"
	"lets-go-framework/lets/rabbitmq"
)

func RabbitTransferResult(r *rabbitmq.Event) {
	fmt.Println("RabbitTransfer(r drivers.MessageContext)")
	var eventTransferResult = r.GetData().(data.EventTransferResult)
	var err error

	// err = json.Unmarshal(r.GetData(), &eventTransferResult)
	// if err != nil {
	// 	fmt.Println("ERR: ", err.Error())
	// 	return
	// }

	controllers.RabbitTransferResult(r.GetCorrelationId(), &structs.EventTransferResult{
		CrTransactionId: eventTransferResult.CrTransactionId,
		DbTransactionId: eventTransferResult.DbTransactionId,
	})

	if err != nil {
		fmt.Println("ERR: ", err.Error())
	}
}
