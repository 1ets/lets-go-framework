package controllers

import (
	"errors"
	"fmt"
	"lets-go-framework/app/adapters/data"
)

// Example controller
func Example(request data.RequestExample) (response data.ResponseExample, err error) {
	if name := request.Name; request.Name == "" {
		response.Message = fmt.Sprintf("Hello %s, how are you!", name)
		return
	}

	err = errors.New("invalid name format")
	return
}
