package structs

type EventTransferResult struct {
	CrTransactionId uint `json:"cr_transaction_id"`
	DbTransactionId uint `json:"db_transaction_id"`
}

type EventBalanceTransferResult struct {
	CrBalance float64 `json:"cr_balance"`
	DbBalance float64 `json:"db_balance"`
}
