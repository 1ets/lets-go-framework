package controllers

import (
	"lets-go-framework/app/adapters"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/structs"
	"lets-go-framework/lets"

	"github.com/gin-gonic/gin"
)

// HTTP Handler for get account information
func HttpTransactionGet(g *gin.Context) {
	var response interface{}
	var err error

	// Get id from uri
	var request structs.HttpAccountRequestRegister
	if err := g.Bind(&request); err != nil {
		lets.HttpResponseJson(g, response, err)
		return
	}

	// Call account service
	svcAccount := adapters.ApiAccount
	response, err = svcAccount.Insert(g, &data.RequestAccountInsert{
		Name:    request.Name,
		Balance: float64(request.Balance),
	})

	lets.HttpResponseJson(g, response, err)
}
