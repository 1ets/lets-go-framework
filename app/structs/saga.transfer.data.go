package structs



type SagaTransferData struct {
	SenderId   uint
	ReceiverId uint
	Amount     float64
	CrTrxId    uint
	DbTrxId    uint
	CrBalance  float64
	DbBalance  float64
}
