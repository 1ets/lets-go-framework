package controllers

import (
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/orchestrator"
	"lets-go-framework/app/structs"
	"lets-go-framework/lets"
	"net/http"
)

// Controller for create synchronus transaction, immediately response with state.
func SagaStateless(request data.RequestTransfer) (response data.ResponseTransfer, err error) {
	var data = structs.SagaTransferData{
		SenderId:   request.SenderId,
		ReceiverId: request.ReceiverId,
		Amount:     request.Amount,
	}

	state, err := orchestrator.SagaTransfer(&data)
	if err != nil {
		lets.LogE(err.Error())
		return
	}

	// lets.LogI("State: %v", state)

	if state == orchestrator.StateNotificationCreated {
		response.Code = http.StatusAccepted
		response.Status = "success"
		response.Message = "Successfully transfer money."

		return
	}

	response.Code = http.StatusNotAcceptable
	response.Status = "failed"
	response.Message = "Transfer was canceled"

	return
}

// Controller for create asynchronus transaction, immediately response without state.
// The scenario is user will receive notification, sms, email or other.
func SagaStateful(request data.RequestTransfer) (response data.ResponseTransfer, err error) {
	var data = structs.SagaTransferData{
		SenderId:   request.SenderId,
		ReceiverId: request.ReceiverId,
		Amount:     request.Amount,
	}

	go func(data *structs.SagaTransferData) {
		state, err := orchestrator.SagaTransfer(data)
		if err != nil {
			lets.LogE(err.Error())
			return
		}

		lets.LogI("State: %v", state)
	}(&data)

	response.Code = http.StatusAccepted
	response.Status = "success"
	response.Message = "Transaction is bein processed"

	return
}
