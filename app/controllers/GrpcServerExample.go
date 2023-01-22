package controllers

import (
	"fmt"
	"lets-go-framework/app/adapters/data"
	"net/http"
)

// gRPC server example controller
func GrpcServerExample(request data.RequestExample) (response data.ResponseExample, err error) {
	response.Code = http.StatusCreated
	response.Status = "success"
	response.Greeting = fmt.Sprintf("Hello %s! this message was created by grpc server.", request.Name)

	return
}
