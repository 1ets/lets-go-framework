package structs

type EventTransferResult struct {
	CrTransactionId uint `json:"cr_transaction_id"`
	DbTransactionId uint `json:"db_transaction_id"`
}
