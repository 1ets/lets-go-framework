package servers

import (
	"fmt"
	"lets-go-framework/adapters/data"
	"lets-go-framework/app/controllers"
	"lets-go-framework/app/structs"
	"lets-go-framework/lets/rabbitmq"
)

func RabbitBalanceResult(r *rabbitmq.Event) {
	fmt.Println("RabbitBalanceResult(r drivers.MessageContext)")
	var eventBalanceTransferResult = r.GetData().(data.EventBalanceTransferResult)
	var err error

	// err = json.Unmarshal(r.GetData(), &eventBalanceTransferResult)
	// if err != nil {
	// 	fmt.Println("ERR: ", err.Error())
	// 	return
	// }

	controllers.RabbitBalanceResult(r.GetCorrelationId(), &structs.EventBalanceTransferResult{
		CrBalance: eventBalanceTransferResult.CrBalance,
		DbBalance: eventBalanceTransferResult.DbBalance,
	})

	if err != nil {
		fmt.Println("ERR: ", err.Error())
	}
}
