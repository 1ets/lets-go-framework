package data

type RequestGetAccount struct {
	Filter *FilterAccount `json:"filter,omitempty"`
}

type FilterAccount struct {
	Id int32 `json:"id,omitempty"`
}

type ResponseGetAccount struct {
	Id      uint    `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}
