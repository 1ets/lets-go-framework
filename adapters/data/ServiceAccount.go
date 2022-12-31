package data

type RequestGetAccount struct {
	Filter *FilterAccount `json:"filter,omitempty"`
}

type FilterAccount struct {
	Id int32 `json:"id,omitempty"`
}
