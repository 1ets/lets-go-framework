package controllers

import (
	"fmt"
	"lets-go-framework/app/adapters/clients"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/models"
	"lets-go-framework/app/repository"
	"lets-go-framework/lets"
	"net/http"
)

// Example controller.
func Example(request data.RequestExample) (response data.ResponseExample, err error) {
	if name := request.Name; request.Name != "" {
		response.Greeting = fmt.Sprintf("Hello %s, how are you!", name)
		return
	}

	response.Code = http.StatusBadRequest
	response.Status = "format_error"

	return
}

// Example controller interaction with database.
func DatabaseExample() (response data.ResponseExample, err error) {
	users := repository.User

	// Repository call
	data, err := users.Get()
	if err != nil {
		return
	}

	// Create output to terminal
	go func() {
		for _, u := range data {
			lets.LogI(u.Name)
		}
	}()

	// Send back response to adapter
	response.Greeting = fmt.Sprintf("We have %v users!", len(data))

	return
}

// Example controller interaction with database.
func RedisExample() (response data.ResponseExample, err error) {
	user := models.User{
		Name:  "John Doe",
		Email: "johndoe@mail.com",
	}

	repository.RedisExample.SaveUser(&user)

	userRedis := repository.RedisExample.GetUser()

	// Send back response to adapter
	response.Greeting = fmt.Sprintf("Hello %s! We test redis.", userRedis.Name)

	return
}

// gRPC server example controller.
func GrpcServerExample(request data.RequestExample) (response data.ResponseExample, err error) {
	response.Code = http.StatusOK
	response.Status = "success"
	response.Greeting = fmt.Sprintf("Hello %s! this message was created by grpc server.", request.Name)

	return
}

// gRPC client example controller.
func GrpcClientExample(request data.RequestExample) (response data.ResponseExample, err error) {
	response, err = clients.GrpcExample.Greeting(&request)
	if err != nil {
		return
	}

	return
}

// RabbitMQ consumer example controller.
func RabbitConsumerExample(request data.RequestExample) (response data.ResponseExample, err error) {
	response.Code = http.StatusOK
	response.Status = "success"
	response.Greeting = fmt.Sprintf("Hello %s! this message was created by rabbit consumer.", request.Name)

	return
}

// RabbitMQ client example controller.
func RabbitPublisherExample(request data.RequestExample, mode string) (response data.ResponseExample, err error) {
	// Wait for reply.
	if mode == "sync" {
		response, err = clients.RabbitExample.GreetingSync(&request)
		if err != nil {
			return
		}
		return
	}

	// No wait reply, just send.
	response, err = clients.RabbitExample.GreetingAsync(&request)
	if err != nil {
		return
	}

	return
}

func CreateSignatureExample() string {
	payload := "Lets Go Framework"

	crypto := lets.Crypto{}
	crypto.SetPrivateKeyFile("keys/private.pem")
	crypto.SetPayloadString(payload)
	crypto.CreateSignatureSHA256WithRSA()

	return crypto.GetSignatureBase64()
}

func VerifySignatureExample(signature string) error {
	payload := "Lets Go Framework"

	crypto := lets.Crypto{}
	crypto.SetPublicKeyFile("keys/public.pem")
	crypto.SetPayloadString(payload)
	crypto.SetSignatureBase64(signature)

	return crypto.VerifySignatureSHA256WithRSA()
}
