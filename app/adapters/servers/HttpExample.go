package servers

import (
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/controllers"
	"lets-go-framework/lets"

	"github.com/gin-gonic/gin"
)

// HTTP Handler for get list of accounts
func HttpPostExample(g *gin.Context) {
	var response data.ResponseExample
	var err error

	// Bind json body into struct format
	var request data.RequestExample
	if err = g.Bind(&request); err != nil {
		lets.HttpResponse(g, response, err)
		return
	}

	// Call example controller
	response, err = controllers.Example(request)
	lets.HttpResponse(g, response, err)
}
