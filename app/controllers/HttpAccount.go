package controllers

import (
	"lets-go-framework/app/adapters"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/structs"
	"lets-go-framework/lets"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTP Handler for get list of accounts
func HttpAccountGet(g *gin.Context) {
	var response interface{}
	var err error

	// Call account service
	svcAccount := adapters.ApiAccount
	response, err = svcAccount.Get(g, &data.RequestAccountGet{})

	lets.HttpResponse(g, response, err)
}

// HTTP Handler for get account information
func HttpAccountFind(g *gin.Context) {
	var response interface{}
	var err error

	// Get id from uri
	var request structs.HttpAccountRequestFind
	if err := g.BindUri(&request); err != nil {
		lets.HttpResponse(g, response, err)
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

	lets.HttpResponse(g, response, err)
}

// HTTP Handler for get account information
func HttpAccountRegister(g *gin.Context) {
	var response interface{}
	var err error

	// Get id from uri
	var request structs.HttpAccountRequestRegister
	if err := g.Bind(&request); err != nil {
		lets.HttpResponse(g, response, err)
		return
	}

	// Call account service
	svcAccount := adapters.ApiAccount
	response, err = svcAccount.Insert(g, &data.RequestAccountInsert{
		Name:    request.Name,
		Balance: float64(request.Balance),
	})

	lets.HttpResponse(g, response, err)
}

// HTTP Handler for get account information
func HttpAccountUpdate(g *gin.Context) {
	var response interface{}
	var err error

	// Get id from uri
	var find structs.HttpAccountRequestFind
	if err := g.BindUri(&find); err != nil {
		lets.HttpResponse(g, response, err)
		return
	}

	// Get id from body
	var update structs.HttpAccountRequestUpdate
	if err := g.Bind(&update); err != nil {
		lets.HttpResponse(g, response, err)
		return
	}

	// Call account service
	svcAccount := adapters.ApiAccount
	response, err = svcAccount.Update(g, &data.RequestAccountUpdate{
		Where: data.AccountUpdateWhere{
			Id: uint(find.Id),
		},
		Data: data.AccountUpdateData{
			Name: update.Name,
		},
	})

	lets.HttpResponse(g, response, err)
}

// HTTP Handler for get account information
func HttpAccountRemove(g *gin.Context) {
	var response interface{}
	var err error

	// Get id from uri
	var find structs.HttpAccountRequestFind
	if err := g.Bind(&find); err != nil {
		lets.HttpResponse(g, response, err)
		return
	}

	// Call account service
	svcAccount := adapters.ApiAccount
	response, err = svcAccount.Delete(g, &data.RequestAccountDelete{
		Id: uint(find.Id),
	})

	lets.HttpResponse(g, response, err)
}
