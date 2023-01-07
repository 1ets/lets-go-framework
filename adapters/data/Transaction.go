package data

type RequestGetTransaction struct {
	Filter string `json:"filter,omitempty"`
}

type FilterTransaction struct {
	Id int32 `json:"id,omitempty"`
}

type EventTransfer struct {
	SenderId   int
	ReceiverId int
	Amount     float32
}

type EventTransferResult struct {
	SenderId   int
	ReceiverId int
	Amount     float32
}
