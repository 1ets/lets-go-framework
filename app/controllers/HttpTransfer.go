package controllers

import (
	"lets-go-framework/app/orchestrator"
	"lets-go-framework/app/structs"
	"lets-go-framework/lets"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kataras/golog"
)

// HTTP Handler for normal transfer
func HttpTransferStateless(g *gin.Context) {
	var response interface{}
	var err error

	// Get id from uri
	var request structs.HttpTransferRequest
	if err := g.Bind(&request); err != nil {
		lets.HttpResponse(g, response, err)
		return
	}

	response = structs.DefaultHttpResponse{
		Code:    http.StatusAccepted,
		Status:  "success",
		Message: "Processing transaction",
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
		lets.HttpResponse(g, response, err)
		return
	}

	lets.HttpResponse(g, response, err)
}

// HTTP Handler for run http transfer in stateful mode
func HttpTransferStatefull(g *gin.Context) {
	var response interface{}
	var err error

	var request structs.HttpTransferRequest
	if err := g.Bind(&request); err != nil {
		lets.HttpResponse(g, response, err)
		return
	}

	response = structs.DefaultHttpResponse{
		Code:    http.StatusAccepted,
		Status:  "success",
		Message: "Processing transaction",
	}

	// Response the request asap
	lets.HttpResponse(g, response, err)

	// Run as a go routines and manage from background
	go func(request structs.HttpTransferRequest) {
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
			return
		}
	}(request)
}
