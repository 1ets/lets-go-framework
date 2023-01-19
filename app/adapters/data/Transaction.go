package data

type RequestGetTransaction struct {
	Filter string `json:"filter,omitempty"`
}

type FilterTransaction struct {
	Id int32 `json:"id,omitempty"`
}
