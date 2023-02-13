package servers

import (
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/lets"

	"github.com/gin-gonic/gin"
)

// HTTP Handler for get list of accounts
func HttpGenerateKey(g *gin.Context) {
	// var request data.RequestExample
	var response data.ResponseExample
	var err error

	// Bind json body into struct format
	// if err = g.Bind(&request); err != nil {
	// 	lets.HttpResponseJson(g, response, err)
	// 	return
	// }

	// Call example controller
	// response, err = controllers.Example(request)

	var c lets.Crypto
	c.GenerateKey(lets.RSA4096)
	lets.LogD(c.GetPrivateKey())
	lets.LogD(c.GetPublicKey())
	c.SavePrivateKey("private.pem")
	c.SavePublicKey("public.pem")

	// Write json response
	lets.HttpResponseJson(g, response, err)
}
