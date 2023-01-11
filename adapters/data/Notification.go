package data

type EventNotification struct {
	SenderId   uint    `json:"id_sender"`
	ReceiverId uint    `json:"id_receiver"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
}
