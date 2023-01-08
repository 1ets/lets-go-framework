package data

type RequestGetTransaction struct {
	Filter string `json:"filter,omitempty"`
}

type FilterTransaction struct {
	Id int32 `json:"id,omitempty"`
}

type EventTransfer struct {
	SenderId   uint    `json:"id_sender"`
	ReceiverId uint    `json:"id_receiver"`
	Amount     float64 `json:"amount"`
}

type EventTransferResult struct {
	Code    uint16 `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
