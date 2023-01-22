package controllers

import (
	"lets-go-framework/app/adapters/clients"
	"lets-go-framework/app/adapters/data"
)

// gRPC client example controller
func GrpcClientExample(request data.RequestExample) (response data.ResponseExample, err error) {
	response, err = clients.GrpcExample.Greeting(&request)
	if err != nil {
		return
	}

	return
}
