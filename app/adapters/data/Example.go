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

type RequestSignatureExample struct {
}

type ResponseSignatureExample struct {
	types.Response

	Signature string `json:"signature"`
}

type RequestVerifySignatureExample struct {
	Signature string `json:"signature"`
}

type ResponsVerifyeSignatureExample struct {
	types.Response

	Result string `json:"result"`
}
