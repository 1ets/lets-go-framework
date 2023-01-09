package data

type EventTransfer struct {
	SenderId   uint    `json:"id_sender"`
	ReceiverId uint    `json:"id_receiver"`
	Amount     float64 `json:"amount"`
}

type EventTransferResult struct {
	CrTransactionId uint `json:"cr_transaction_id"`
	DbTransactionId uint `json:"db_transaction_id"`
}
