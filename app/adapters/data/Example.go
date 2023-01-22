package data

import "lets-go-framework/lets/types"

type RequestExample struct {
	Name string `json:"name"`
}

type ResponseExample struct {
	types.Response

	Greeting string `json:"greeting"`
}

type RequestDatabaseExample struct{}

type ResponsDatabaseeExample struct {
	types.Response

	Note string `json:"note"`
}
