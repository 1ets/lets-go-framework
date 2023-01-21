package controllers

import (
	"fmt"
	"lets-go-framework/app/adapters/data"
	"net/http"
)

// Example controller
func Example(request data.RequestExample) (response data.ResponseExample, err error) {
	if name := request.Name; request.Name != "" {
		response.Greeting = fmt.Sprintf("Hello %s, how are you!", name)
		return
	}

	response.Code = http.StatusBadRequest
	response.Status = "format_error"

	return
}
