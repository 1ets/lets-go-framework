package servers

import (
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/controllers"
	"lets-go-framework/lets"

	"github.com/gin-gonic/gin"
)

func HttpGetSignatureExample(g *gin.Context) {
	var response data.ResponseSignatureExample
	var err error

	// Call example controller
	response.Signature = controllers.CreateSignatureExample()

	// Write json response
	lets.HttpResponseJson(g, response, err)
}

func HttpVerifySignatureExample(g *gin.Context) {
	var request data.RequestVerifySignatureExample
	var response data.ResponsVerifyeSignatureExample
	var err error

	response.Result = "Verified"

	// Bind json body into struct format
	if err = g.Bind(&request); err != nil {
		lets.HttpResponseJson(g, response, err)
		return
	}

	// Call example controller
	err = controllers.VerifySignatureExample(request.Signature)
	if err != nil {
		response.Result = "Unverified"
	}

	// Write json response
	lets.HttpResponseJson(g, response, err)
}

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

// HTTP Handler for get list of accounts
func HttpEncryptDecrypt(g *gin.Context) {
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

	payload := "Lets Go Framework"
	lets.LogD("Payload: %s", payload)

	var c lets.Crypto
	c.SetPrivateKeyFile("keys/private.pem")
	c.SetPublicKeyFile("keys/public.pem")

	_, b64 := c.EncryptOAEP(payload)
	lets.LogD("Encrypted: %s", b64)

	decrypted := c.DecryptB64OAEP(b64)
	lets.LogD("Decrypted: %s", decrypted)

	// Write json response
	lets.HttpResponseJson(g, response, err)
}
