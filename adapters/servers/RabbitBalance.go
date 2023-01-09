package servers

import (
	"encoding/json"
	"fmt"
	"lets-go-framework/adapters/data"
	"lets-go-framework/app/controllers"
	"lets-go-framework/app/structs"
	"lets-go-framework/lets/drivers"
)

func RabbitBalanceResult(r drivers.MessageContext) {
	fmt.Println("RabbitBalanceResult(r drivers.MessageContext)")
	var eventTransferResult data.EventTransferResult
	var err error

	err = json.Unmarshal(r.GetData(), &eventTransferResult)
	if err != nil {
		fmt.Println("ERR: ", err.Error())
		return
	}

	controllers.RabbitBalanceResult(r.GetCorrelationId(), &structs.EventBalanceTransferResult{})

	if err != nil {
		fmt.Println("ERR: ", err.Error())
	}
}
