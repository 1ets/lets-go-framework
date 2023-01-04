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
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type ResponseAccountInsert struct {
	Code   uint16 `json:"code"`
	Status string `json:"status"`
}

type RequestAccountUpdate struct {
	Where AccountUpdateWhere `json:"where"`
	Data  AccountUpdateData  `json:"data"`
}

type AccountUpdateWhere struct {
	Id uint `json:"id"`
}

type AccountUpdateData struct {
	Name string `json:"name,omitempty"`
}

type ResponseAccountUpdate struct {
	Code   uint16 `json:"code"`
	Status string `json:"status"`
}

type RequestAccountDelete struct {
	Id uint `json:"id"`
}

type ResponseAccountDelete struct {
	Code   uint16 `json:"code"`
	Status string `json:"status"`
}
