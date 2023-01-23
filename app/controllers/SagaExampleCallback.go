package controllers

import (
	"lets-go-framework/app/adapters/data"
	"net/http"
)

// Simulation for transfer service provider.
func RabbitSagaTransferServiceProvider(request data.RequestTransfer) (response data.ResponseTransfer, err error) {

	response.Code = http.StatusCreated
	response.Status = "success"
	response.CrTransactionId = 1
	response.DbTransactionId = 2

	return
}

// Simulation for balance service provider.
func RabbitSagaBalanceServiceProvider(request data.RequestBalance) (response data.ResponseBalance, err error) {

	// Success
	response.Code = http.StatusOK
	response.Status = "success"
	response.CrBalance = 1000
	response.DbBalance = 2000

	// Fails
	// response.Code = http.StatusOK
	// response.Status = "success"
	// response.CrBalance = 1000
	// response.DbBalance = 2000

	return
}

// Simulation for notification service provider.
func RabbitSagaNotificationServiceProvider(request data.RequestNotification) (response data.ResponseNotification, err error) {

	// Success
	response.Code = http.StatusOK
	response.Status = "success"

	// Fails
	// response.Code = http.StatusInternalServerError
	// response.Status = "error"

	return
}
