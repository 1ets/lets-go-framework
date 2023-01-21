package data

import "lets-go-framework/lets/types"

type RequestExample struct {
	Name string `json:"name"`
}

type ResponseExample struct {
	types.Response

	Greeting string `json:"greeting"`
}
