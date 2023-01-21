package controllers

import (
	"fmt"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/repository"
	"lets-go-framework/lets"
)

// Example controller get users
func DatabaseExample() (response data.ResponseExample, err error) {
	users := repository.User

	// Repository call
	data, err := users.Get()
	if err != nil {
		return
	}

	// Create output to terminal
	for _, u := range data {
		lets.LogI(u.Name)
	}

	// Send back response to adapter
	response.Greeting = fmt.Sprintf("We have %v users!", len(data))

	return
}
