package orchestrator

import (
	"context"
	"fmt"
	"lets-go-framework/app/adapters/clients"
	"lets-go-framework/app/adapters/data"
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

type Transfer struct {
	bucket *DataTransferEvent
}

var WaitTransfer = make(map[string](chan structs.EventTransferResult))

func (t *Transfer) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	fmt.Println("Transfer")

	var correlationId = fmt.Sprintf("%v", time.Now().Unix())
	WaitTransfer[correlationId] = make(chan structs.EventTransferResult)
	defer delete(WaitTransfer, correlationId)

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

	var data = <-WaitTransfer[correlationId]
	t.bucket.CrTrxId = data.CrTransactionId
	t.bucket.DbTrxId = data.DbTransactionId

	return &ActivityTransfer{
		state: StateTransferSucceed,
	}, nil
}

// Rollback state
type TransferRollback struct {
	bucket *DataTransferEvent
}

func (t *TransferRollback) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	fmt.Println("TransferRollback")

	// Call transaction service
	svcTransaction := clients.RabbitTransfer
	err := svcTransaction.TransferRollback(&data.EventTransferRollback{
		CrTransactionId: t.bucket.CrTrxId,
		DbTransactionId: t.bucket.DbTrxId,
	})

	return &ActivityTransfer{
		state: StateTransferCanceled,
	}, err
}

type BalanceTransfer struct {
	bucket *DataTransferEvent
}

// Make it public so can accessed by controller handler
var WaitBalanceTransfer = make(map[string](chan structs.EventBalanceTransferResult))

func (t *BalanceTransfer) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	fmt.Println("BalanceCommitStart")

	var correlationId = fmt.Sprintf("%v", time.Now().Unix())
	WaitBalanceTransfer[correlationId] = make(chan structs.EventBalanceTransferResult)
	defer delete(WaitBalanceTransfer, correlationId)

	// Call account service
	svcBalance := clients.RabbitBalance
	err := svcBalance.BalanceTransfer(correlationId, &data.EventTransfer{
		SenderId:   t.bucket.SenderId,
		ReceiverId: t.bucket.ReceiverId,
		Amount:     t.bucket.Amount,
	})

	if err != nil {
		return &ActivityTransfer{
			state: StateTransferFailed,
		}, err
	}

	var response = <-WaitBalanceTransfer[correlationId]
	t.bucket.CrBalance = response.CrBalance
	t.bucket.DbBalance = response.DbBalance

	// fmt.Printf("if %f	<= 0 || %f <= 0 {", response.CrBalance, response.DbBalance)

	if response.CrBalance <= 0 || response.DbBalance <= 0 {
		err := svcBalance.BalanceRollback(&data.EventTransfer{
			SenderId:   t.bucket.SenderId,
			ReceiverId: t.bucket.ReceiverId,
			Amount:     t.bucket.Amount,
		})

		fmt.Println(err)

		return &ActivityTransfer{
			state: StateBalanceTranferFailed,
		}, nil
	}

	return &ActivityTransfer{
		state: StateBalanceTranferSucceed,
	}, nil
}

type BalanceRollback struct {
	bucket *DataTransferEvent
}

func (t *BalanceRollback) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	fmt.Println("BalanceRollback")

	// Call account service
	svcBalance := clients.RabbitBalance
	err := svcBalance.BalanceRollback(&data.EventTransfer{
		SenderId:   t.bucket.SenderId,
		ReceiverId: t.bucket.ReceiverId,
		Amount:     t.bucket.Amount,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	return &ActivityTransfer{
		state: StateBalanceTranferCanceled,
	}, nil
}

type NotifyUserStart struct {
	bucket *DataTransferEvent
}

func (a *NotifyUserStart) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	fmt.Println("NotifyUserStart")

	// Call account service
	svcNotification := clients.RabbitNotification
	err := svcNotification.Notify(&data.EventNotification{
		SenderId:   a.bucket.SenderId,
		ReceiverId: a.bucket.ReceiverId,
		Amount:     a.bucket.Amount,
		Status:     "SUCCESS",
	})

	if err != nil {
		fmt.Println(err.Error())
	}

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
func SagaTransfer(data *structs.HttpTransferRequest) (state saga.State, err error) {
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
				Aggregator: &Transfer{
					bucket: &bucket,
				},
				SuccessState: StateTransferSucceed,
				FailureState: StateTransferFailed,
				Compensation: &TransferRollback{
					bucket: &bucket,
				},
				RolledBackState: StateTransferCanceled,
			},
			{
				Aggregator: &BalanceTransfer{
					bucket: &bucket,
				},
				SuccessState: StateBalanceTranferSucceed,
				FailureState: StateBalanceTranferFailed,
				Compensation: &BalanceRollback{
					bucket: &bucket,
				},
				RolledBackState: StateBalanceTranferCanceled,
			},
			{
				Aggregator: &NotifyUserStart{
					bucket: &bucket,
				},
				SuccessState:    StateNotifyUserSucceed,
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

	state = finalTx.State()

	return
}
