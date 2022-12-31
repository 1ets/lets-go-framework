package controllers

import (
	"lets-go-framework/adapters"
	"lets-go-framework/adapters/data"
	"lets-go-framework/lets"

	"github.com/gin-gonic/gin"
)

func HttpGetAccount(g *gin.Context) {
	var response interface{}
	var err error

	response, err = adapters.ApiAccount.Get(g, data.RequestGetAccount{
		Filter: &data.FilterAccount{
			Id: 1,
		},
	})

	lets.Response(g, response, err)
}
