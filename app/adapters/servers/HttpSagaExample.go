package servers

import (
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/controllers"
	"lets-go-framework/lets"

	"github.com/gin-gonic/gin"
)

// HTTP Handler for running saga transaction stateless.
func HttpSagaStatelessExample(g *gin.Context) {
	var request data.RequestTransfer
	var response data.ResponseTransfer
	var err error

	// Bind json body into struct format
	if err = g.Bind(&request); err != nil {
		lets.HttpResponseJson(g, response, err)
		return
	}

	// Call example controller
	response, err = controllers.SagaStateless(request)

	// Write json response
	lets.HttpResponseJson(g, response, err)
}

// // HTTP Handler for running saga transaction stateful.
// func HttpSagaStatefullExample(g *gin.Context) {
// 	var request data.RequestExample
// 	var response data.ResponseExample
// 	var err error

// 	// Bind json body into struct format
// 	if err = g.Bind(&request); err != nil {
// 		lets.HttpResponseJson(g, response, err)
// 		return
// 	}

// 	// Call example controller
// 	response, err = controllers.SagaStatefull(request)

// 	// Write json response
// 	lets.HttpResponseJson(g, response, err)
// }
