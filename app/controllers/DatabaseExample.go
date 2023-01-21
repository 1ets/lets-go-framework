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

	data, err := users.Get()
	if err != nil {
		return
	}

	for _, u := range data {
		lets.LogI(u.Name)
	}

	response.Greeting = fmt.Sprintf("We have %v users!", len(data))

	return
}
