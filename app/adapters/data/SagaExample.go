package data

import "lets-go-framework/lets/types"

// Request transfer.
type RequestTransfer struct {
	SenderId   uint    `json:"id_sender" form:"id_sender"`
	ReceiverId uint    `json:"id_receiver" form:"id_receiver"`
	Amount     float64 `json:"amount" form:"amount"`
}

// Response Transfer
type ResponseTransfer struct {
	types.Response

	CrTransactionId uint `json:"cr_transaction_id,omitempty"`
	DbTransactionId uint `json:"db_transaction_id,omitempty"`
}

// Request rollback the transfer.
type RequestTransferRollback struct {
	CrTransactionId uint `json:"cr_transaction_id"`
	DbTransactionId uint `json:"db_transaction_id"`
}

// Response rollback the transfer.
type ResponseTransferRollback struct {
	types.Response

	CrTransactionId uint `json:"cr_transaction_id"`
	DbTransactionId uint `json:"db_transaction_id"`
}

// Request balance.
type RequestBalance struct {
	SenderId   uint    `json:"id_sender"`
	ReceiverId uint    `json:"id_receiver"`
	Amount     float64 `json:"amount"`
}

// Response balance.
type ResponseBalance struct {
	types.Response

	CrBalance float64 `json:"cr_balance"`
	DbBalance float64 `json:"db_balance"`
}

type RequestNotification struct {
	SenderId   uint    `json:"id_sender"`
	ReceiverId uint    `json:"id_receiver"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
}

type ResponseNotification struct {
	types.Response
}
