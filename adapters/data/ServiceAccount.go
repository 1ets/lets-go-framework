package data

type Account struct {
	Id      uint    `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type RequestAccountGet struct{}
type ResponseAccountGet []Account

type RequestAccountFind struct {
	Id uint `json:"id"`
}
type ResponseAccountFind Account

type RequestAccountInsert struct {
	Id      uint    `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Balance float64 `json:"balance,omitempty"`
}
type ResponseAccountInsert struct {
	Code   uint16 `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
}
