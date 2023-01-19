package data

type EventBalanceTransfer struct {
	SenderId   uint    `json:"id_sender"`
	ReceiverId uint    `json:"id_receiver"`
	Amount     float64 `json:"amount"`
}

type EventBalanceTransferResult struct {
	CrBalance float64 `json:"cr_balance"`
	DbBalance float64 `json:"db_balance"`
}
