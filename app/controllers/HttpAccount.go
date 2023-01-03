package controllers

import (
	"lets-go-framework/adapters"
	"lets-go-framework/adapters/data"
	"lets-go-framework/app/structs"
	"lets-go-framework/lets"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTP Handler for get list of accounts
func HttpAccountGet(g *gin.Context) {
	var response interface{}
	var err error
	svcAccount := adapters.ApiAccount

	response, err = svcAccount.Get(g, &data.RequestAccountGet{})

	lets.Response(g, response, err)
}

// HTTP Handler for get account information
func HttpAccountFind(g *gin.Context) {
	var response interface{}
	var err error

	// Get id from uri
	var request structs.HttpAccountRequestFind
	if err := g.BindUri(&request); err != nil {
		lets.Response(g, response, err)
		return
	}

	// Call account service
	svcAccount := adapters.ApiAccount
	dataResponse, err := svcAccount.Find(g, &data.RequestAccountFind{
		Id: uint(request.Id),
	})

	// Set response
	response = dataResponse

	// Intercept response
	if dataResponse.Id == 0 {
		response = &structs.HttpAccountResponseDefault{
			Code:    http.StatusNotFound,
			Status:  "not_found",
			Message: "Record not found",
		}
	}

	lets.Response(g, response, err)
}

// HTTP Handler for get account information
func HttpAccountRegister(g *gin.Context) {
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
