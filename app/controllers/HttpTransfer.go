package controllers

import (
	"lets-go-framework/adapters"
	"lets-go-framework/adapters/clients"
	"lets-go-framework/adapters/data"
	"lets-go-framework/app/structs"
	"lets-go-framework/lets"

	"github.com/gin-gonic/gin"
)

// HTTP Handler for get account information
func HttpTransferSuccess(g *gin.Context) {
	var response interface{}
	var err error

	// Get id from uri
	var request structs.HttpTransferRequest
	if err := g.Bind(&request); err != nil {
		lets.Response(g, response, err)
		return
	}

	// Call account service
	svcTransaction := clients.RabbitTransfer
	err = svcTransaction.Transfer(&data.EventTransfer{
		SenderId:   request.SenderId,
		ReceiverId: request.ReceiverId,
		Amount:     request.Amount,
	})

	lets.Response(g, response, err)
}

// HTTP Handler for get account information
func HttpTransferFailed(g *gin.Context) {
	var response interface{}
	var err error

	// Get id from uri
	var request structs.HttpAccountRequestRegister
	if err := g.Bind(&request); err != nil {
		lets.Response(g, response, err)
		return
	}

	// Call account service
	svcAccount := adapters.ApiAccount
	response, err = svcAccount.Insert(g, &data.RequestAccountInsert{
		Name:    request.Name,
		Balance: float64(request.Balance),
	})

	lets.Response(g, response, err)
}
