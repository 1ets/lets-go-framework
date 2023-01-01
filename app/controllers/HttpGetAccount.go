package controllers

import (
	"lets-go-framework/adapters"
	"lets-go-framework/adapters/data"
	"lets-go-framework/lets"

	"github.com/gin-gonic/gin"
)

// HTTP Controller Handler for get account information
func HttpGetAccount(g *gin.Context) {
	var response interface{}
	var err error
	svcAccount := adapters.ApiAccount

	response, err = svcAccount.Get(g, &data.RequestGetAccount{
		Filter: &data.FilterAccount{
			Id: 1,
		},
	})

	lets.Response(g, response, err)
}
