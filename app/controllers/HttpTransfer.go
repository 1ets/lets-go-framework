package controllers

import (
	"lets-go-framework/adapters"
	"lets-go-framework/adapters/data"
	"lets-go-framework/app/orchestrator"
	"lets-go-framework/app/structs"
	"lets-go-framework/lets"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
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

	state, err := orchestrator.SagaTransfer(&request)
	if err != nil {
		golog.Error(err.Error())
		return
	}

	golog.Infof("State: %v", state)

	if state == orchestrator.StateTransferCanceled {
		response = structs.DefaultHttpResponse{
			Code:    http.StatusNotAcceptable,
			Status:  "failed",
			Message: "Transfer was canceled",
		}
		lets.Response(g, response, err)
		return
	}

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
