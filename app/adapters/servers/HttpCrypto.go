package servers

import (
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/controllers"
	"lets-go-framework/lets"

	"github.com/gin-gonic/gin"
)

// HTTP Handler for get list of accounts
func HttpGenerateKey(g *gin.Context) {
	var response data.ResponseExample
	var err error

	// Call example controller
	response, err = controllers.GenerateKeys()

	// Write json response
	lets.HttpResponseJson(g, response, err)
}

func HttpGetSignatureExample(g *gin.Context) {
	var response data.ResponseSignatureExample
	var err error

	// Call example controller
	response.Signature, err = controllers.CreateSignature()

	// Write json response
	lets.HttpResponseJson(g, response, err)
}

func HttpVerifySignatureExample(g *gin.Context) {
	var request data.RequestVerifySignatureExample
	var response data.ResponsVerifyeSignatureExample
	var err error

	// Bind json body into struct format
	if err = g.Bind(&request); err != nil {
		lets.HttpResponseJson(g, response, err)
		return
	}

	// Call example controller
	response.Result, err = controllers.VerifySignature(request.Signature)

	// Write json response
	lets.HttpResponseJson(g, response, err)
}

// HTTP Handler for encryption decryption demo.
func HttpEncryptDecrypt(g *gin.Context) {
	var response data.ResponseExample

	err := controllers.EncryptDecrypt()

	// Write json response
	lets.HttpResponseJson(g, response, err)
}
