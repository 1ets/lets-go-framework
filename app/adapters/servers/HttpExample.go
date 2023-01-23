package servers

import (
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/controllers"
	"lets-go-framework/lets"

	"github.com/gin-gonic/gin"
)

// HTTP Handler for get list of accounts
func HttpPostExample(g *gin.Context) {
	var request data.RequestExample
	var response data.ResponseExample
	var err error

	// Bind json body into struct format
	if err = g.Bind(&request); err != nil {
		lets.HttpResponseJson(g, response, err)
		return
	}

	// Call example controller
	response, err = controllers.Example(request)

	// Write json response
	lets.HttpResponseJson(g, response, err)
}

// HTTP Handler for get list of users
func HttpGetDatabaseExample(g *gin.Context) {
	// var request data.RequestExample
	var response data.ResponseExample
	var err error

	// Call example controller
	response, err = controllers.DatabaseExample()

	// Write json response
	lets.HttpResponseJson(g, response, err)
}

// HTTP Handler for get list of users
func HttpGrpcExample(g *gin.Context) {
	var request data.RequestExample
	var response data.ResponseExample
	var err error

	// Bind json body into struct format
	if err = g.Bind(&request); err != nil {
		lets.HttpResponseJson(g, response, err)
		return
	}

	// Call example controller
	response, err = controllers.GrpcClientExample(request)

	// Write json response
	lets.HttpResponseJson(g, response, err)
}

// HTTP Handler for get list of users
func HttpRabbitAsyncExample(g *gin.Context) {
	var request data.RequestExample
	var response data.ResponseExample
	var err error

	// Bind json body into struct format
	if err = g.Bind(&request); err != nil {
		lets.HttpResponseJson(g, response, err)
		return
	}

	// Call example controller
	response, err = controllers.RabbitPublisherExample(request, "async")

	// Write json response
	lets.HttpResponseJson(g, response, err)
}

// HTTP Handler for get list of users
func HttpRabbitSyncExample(g *gin.Context) {
	var request data.RequestExample
	var response data.ResponseExample
	var err error

	// Bind json body into struct format
	if err = g.Bind(&request); err != nil {
		lets.HttpResponseJson(g, response, err)
		return
	}

	// Call example controller
	response, err = controllers.RabbitPublisherExample(request, "sync")

	// Write json response
	lets.HttpResponseJson(g, response, err)
}
