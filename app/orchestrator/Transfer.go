package orchestrator

import (
	"context"
	"fmt"
	"lets-go-framework/adapters/clients"
	"lets-go-framework/adapters/data"
	"lets-go-framework/app/structs"
	"time"

	"github.com/bongnv/saga"
)

const (
	StateTransferInit saga.State = iota
	StateTransferSucceed
	StateTransferFailed
	StateTransferCanceled
	StateBalanceTranferSucceed
	StateBalanceTranferFailed
	StateBalanceTranferCanceled
	StateNotifyUserInit
	StateNotifyUserSucceed
	StateNotifyUserFailed
	StateNotifyUserRetry
)

type ActivityTransfer struct {
	state saga.State
}

func (t *ActivityTransfer) State() saga.State {
	return t.state
}

type TransferStart struct {
	bucket *DataTransferEvent
}

var WaitTransferStart = make(map[string](chan structs.EventTransferResult))

func (t *TransferStart) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	fmt.Println("TransferStart")

	var correlationId = fmt.Sprintf("%v", time.Now().Unix())
	WaitTransferStart[correlationId] = make(chan structs.EventTransferResult)
	defer delete(WaitTransferStart, correlationId)

	// Call account service
	svcTransaction := clients.RabbitTransfer
	err := svcTransaction.Transfer(correlationId, &data.EventTransfer{
		SenderId:   t.bucket.SenderId,
		ReceiverId: t.bucket.ReceiverId,
		Amount:     t.bucket.Amount,
	})

	if err != nil {
		return &ActivityTransfer{
			state: StateTransferFailed,
		}, err
	}

	var data = <-WaitTransferStart[correlationId]
	t.bucket.CrTrxId = data.CrTransactionId
	t.bucket.DbTrxId = data.DbTransactionId

	return &ActivityTransfer{
		state: StateTransferSucceed,
	}, nil
}

type TransferCancel struct {
	bucket *DataTransferEvent
}

func (a *TransferCancel) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	fmt.Println("TransferCancel")
	// Publish to

	return &ActivityTransfer{
		state: StateTransferCanceled,
	}, nil
}

type BalanceTransferStart struct {
	bucket *DataTransferEvent
}

// Make it public so can accessed by controller handler
var WaitBalanceTransfer = make(map[string](chan structs.EventBalanceTransferResult))

func (t *BalanceTransferStart) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	fmt.Println("BalanceCommitStart")

	var correlationId = fmt.Sprintf("%v", time.Now().Unix())
	WaitBalanceTransfer[correlationId] = make(chan structs.EventBalanceTransferResult)
	defer delete(WaitBalanceTransfer, correlationId)

	// Call account service
	svcTraBalance := clients.RabbitBalance
	err := svcTraBalance.BalanceTransfer(correlationId, &data.EventTransfer{
		SenderId:   t.bucket.SenderId,
		ReceiverId: t.bucket.ReceiverId,
		Amount:     t.bucket.Amount,
	})

	if err != nil {
		return &ActivityTransfer{
			state: StateTransferFailed,
		}, err
	}

	var data = <-WaitBalanceTransfer[correlationId]
	t.bucket.CrBalance = data.CrBalance
	t.bucket.DbBalance = data.DbBalance

	if data.CrBalance <= 0 || data.DbBalance <= 0 {
		return &ActivityTransfer{
			state: StateBalanceTranferFailed,
		}, nil
	}

	return &ActivityTransfer{
		state: StateBalanceTranferSucceed,
	}, nil
}

type BalanceTransferCancel struct {
	bucket *DataTransferEvent
}

func (a *BalanceTransferCancel) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	fmt.Println("BalanceCommitCancel")
	// Publish to

	return &ActivityTransfer{
		state: StateTransferCanceled,
	}, nil
}

type NotifyUserStart struct{}

func (a *NotifyUserStart) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	fmt.Println("NotifyUserStart")
	// Publish to

	return &ActivityTransfer{
		state: StateNotifyUserSucceed,
	}, nil
}

type DataTransferEvent struct {
	SenderId   uint
	ReceiverId uint
	Amount     float64
	CrTrxId    uint
	DbTrxId    uint
	CrBalance  float64
	DbBalance  float64
}

// Define saga transfer
func SagaTransfer(data *structs.HttpTransferRequest) (state int, err error) {
	// Create data bucket
	bucket := DataTransferEvent{
		SenderId:   data.SenderId,
		ReceiverId: data.ReceiverId,
		Amount:     data.Amount,
	}

	exampleConfig := saga.Config{
		InitState: StateTransferInit,
		Activities: []saga.Activity{
			{
				Aggregator: &TransferStart{
					bucket: &bucket,
				},
				SuccessState: StateTransferSucceed,
				FailureState: StateTransferFailed,
				Compensation: &TransferCancel{
					bucket: &bucket,
				},
				RolledBackState: StateTransferCanceled,
			},
			{
				Aggregator: &BalanceTransferStart{
					bucket: &bucket,
				},
				SuccessState: StateBalanceTranferSucceed,
				FailureState: StateBalanceTranferFailed,
				Compensation: &BalanceTransferCancel{
					bucket: &bucket,
				},
				RolledBackState: StateBalanceTranferCanceled,
			},
			{
				Aggregator:      &NotifyUserStart{},
				SuccessState:    StateNotifyUserInit,
				FailureState:    StateNotifyUserFailed,
				RolledBackState: StateNotifyUserRetry,
			},
		},
	}

	executor, err := saga.NewExecutor(exampleConfig)
	if err != nil {
		fmt.Printf("Failed to create saga executor, err: %v.\n", err)
		return
	}

	finalTx, err := executor.Execute(context.Background(), &ActivityTransfer{
		state: StateTransferInit,
	})
	if err != nil {
		fmt.Printf("Failed to execute saga transaction, err: %v.\n", err)
		return
	}

	if finalTx.State() == StateTransferCanceled {
		fmt.Println("OK")
	}

	return
}
