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
	StateBalanceCommitSucceed
	StateBalanceCommitFailed
	StateBalanceCommitCanceled
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
	var correlationId = fmt.Sprintf("%v", time.Now().Unix())
	WaitTransferStart[correlationId] = make(chan structs.EventTransferResult)

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

	delete(WaitTransferStart, correlationId)

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

type BalanceCommitStart struct {
	bucket *DataTransferEvent
}

// Make it public so can accessed by controller handler
var WaitBalanceCommitStart = make(map[string](chan structs.EventTransferResult))

func (t *BalanceCommitStart) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	fmt.Println("BalanceCommitStart")

	var correlationId = fmt.Sprintf("%v", time.Now().Unix())
	WaitBalanceCommitStart[correlationId] = make(chan structs.EventTransferResult)

	// Call account service
	svcTransaction := clients.RabbitTransfer
	err := svcTransaction.BalanceCommit(correlationId, &data.EventTransfer{
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

	delete(WaitTransferStart, correlationId)

	return &ActivityTransfer{
		state: StateBalanceCommitSucceed,
	}, nil
}

type BalanceCommitCancel struct {
	bucket *DataTransferEvent
}

func (a *BalanceCommitCancel) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
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
				Aggregator: &BalanceCommitStart{
					bucket: &bucket,
				},
				SuccessState: StateBalanceCommitSucceed,
				FailureState: StateBalanceCommitFailed,
				Compensation: &BalanceCommitCancel{
					bucket: &bucket,
				},
				RolledBackState: StateBalanceCommitCanceled,
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
