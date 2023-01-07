package structs

type HttpTransferRequest struct {
	SenderId   int     `json:"id_sender"`
	ReceiverId int     `json:"id_receiver"`
	Amount     float64 `json:"amount"`
}
