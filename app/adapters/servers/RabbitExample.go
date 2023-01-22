package servers

import (
	"fmt"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/controllers"
	"lets-go-framework/lets"
	"lets-go-framework/lets/types"
)

func RabbitExample(r *types.Event) {
	var request data.RequestExample
	var err error

	lets.Bind(r.GetData(), &request)

	// Call controller
	response, err := controllers.RabbitConsumerExample(request)
	if err != nil {
		lets.LogE("gRPC Server: GrpcExample.Greeting:", err.Error())
		return
	}

	if err != nil {
		fmt.Println("ERR: ", err.Error())
	}

	lets.LogD(lets.ToJson(response))
}
