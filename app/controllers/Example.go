package controllers

import (
	"fmt"
	"lets-go-framework/app/adapters/clients"
	"lets-go-framework/app/adapters/data"
	"net/http"
)

// Example controller
func Example(request data.RequestExample) (response data.ResponseExample, err error) {
	if name := request.Name; request.Name != "" {
		response.Greeting = fmt.Sprintf("Hello %s, how are you!", name)
		return
	}

	response.Code = http.StatusBadRequest
	response.Status = "format_error"

	return
}

// gRPC server example controller
func GrpcServerExample(request data.RequestExample) (response data.ResponseExample, err error) {
	response.Code = http.StatusOK
	response.Status = "success"
	response.Greeting = fmt.Sprintf("Hello %s! this message was created by grpc server.", request.Name)

	return
}

// gRPC client example controller
func GrpcClientExample(request data.RequestExample) (response data.ResponseExample, err error) {
	response, err = clients.GrpcExample.Greeting(&request)
	if err != nil {
		return
	}

	return
}

// RabbitMQ consumer example controller
func RabbitConsumerExample(request data.RequestExample) (response data.ResponseExample, err error) {
	response.Code = http.StatusOK
	response.Status = "success"
	response.Greeting = fmt.Sprintf("Hello %s! this message was created by rabbit consumer.", request.Name)

	return
}

// RabbitMQ client example controller
func RabbitPublisherExample(mode string, request data.RequestExample) (response data.ResponseExample, err error) {
	if mode == "sync" {
		response, err = clients.RabbitExample.GreetingSync(&request)
		if err != nil {
			return
		}
		return
	}

	response, err = clients.RabbitExample.GreetingAsync(&request)
	if err != nil {
		return
	}

	return
}
