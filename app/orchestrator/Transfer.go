package orchestrator

import (
	"context"
	"fmt"
	"lets-go-framework/app/adapters/clients"
	"lets-go-framework/app/adapters/data"
	"lets-go-framework/app/structs"
	"lets-go-framework/lets"
	"net/http"

	"github.com/bongnv/saga"
)

// This scenario is made to make it easier for you to understand how
// the transfer saga works, even though not all processes will be
// implemented by you in reality.
// How it works:
// 1. Initially request a transfer to the transfer service.
// 2. Begin to update the balance to the balance service.
// 3. Send notifications to users.

const (
	StateTransferInit saga.State = iota
	StateTransferSucceed
	StateTransferFailed
	StateTransferCanceled
	StateBalanceTranferSucceed
	StateBalanceTranferFailed
	StateBalanceTranferCanceled
	StateNotificationCreated
	StateNotificationFailed
	StateNotificationRetry
)

type ExampleActivityTransfer struct {
	state saga.State
}

func (t *ExampleActivityTransfer) State() saga.State {
	return t.state
}

// Saga transfer handler.
type ExampleTransfer struct {
	bucket *structs.SagaTransferData
}

// Executor for requesting transaction.
func (t *ExampleTransfer) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	lets.LogI("Saga: Transfer: Starting ...")

	var request = data.RequestTransfer{
		SenderId:   t.bucket.SenderId,
		ReceiverId: t.bucket.ReceiverId,
		Amount:     t.bucket.Amount,
	}

	// Call transfer service
	var response data.ResponseTransfer
	var err error
	if response, err = clients.RabbitSagaExample.Transfer(&request); err != nil {
		lets.LogE("Saga: Transfer: %s", err.Error())
	}

	// Success response
	if response.Code == http.StatusCreated {
		lets.LogI("Saga: Transfer: Success")

		t.bucket.CrTrxId = response.CrTransactionId
		t.bucket.DbTrxId = response.DbTransactionId

		return &ExampleActivityTransfer{
			state: StateTransferSucceed,
		}, nil
	}

	// Error response
	lets.LogI("Saga: Transfer: Failed")
	return &ExampleActivityTransfer{
		state: StateTransferFailed,
	}, err
}

// Saga rollback transfer handler.
type TransferRollback struct {
	bucket *structs.SagaTransferData
}

// Executor for requesting transaction rollback.
func (t *TransferRollback) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	lets.LogI("Saga: TransferRollback: Starting ...")

	var request = data.RequestTransferRollback{
		CrTransactionId: t.bucket.CrTrxId,
		DbTransactionId: t.bucket.DbTrxId,
	}

	// Call transfer rollback service
	var err error
	if _, err = clients.RabbitSagaExample.TransferRollback(&request); err != nil {
		lets.LogE("Saga: Transfer: %s", err.Error())
	}

	lets.LogI("Saga: TransferRollback: End")

	// With confidence assume that the rollback is handled by the service properly,
	// so there is no need to check. If needed, you can check rollback response,
	// with the creativity of your code.
	return &ExampleActivityTransfer{
		state: StateTransferCanceled,
	}, err
}

// Saga balance transfer service handler.
type BalanceTransfer struct {
	bucket *structs.SagaTransferData
}

// Executor for requesting balance transfer.
func (t *BalanceTransfer) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	lets.LogI("Saga: BalanceTransfer: Starting ...")

	var request = data.RequestBalance{
		SenderId:   t.bucket.SenderId,
		ReceiverId: t.bucket.ReceiverId,
		Amount:     t.bucket.Amount,
	}

	// Call transfer service
	var response data.ResponseBalance
	var err error
	if response, err = clients.RabbitSagaExample.Balance(&request); err != nil {
		lets.LogE("Saga: Transfer: %s", err.Error())
	}
	// Success response
	if response.Code == http.StatusOK {
		lets.LogI("Saga: BalanceTransfer: Success")

		t.bucket.CrBalance = response.CrBalance
		t.bucket.DbBalance = response.DbBalance

		return &ExampleActivityTransfer{
			state: StateBalanceTranferSucceed,
		}, nil
	}

	// In this scenario it's a side check in the saga, but that doesn't mean it's a good thing.
	if response.CrBalance <= 0 || response.DbBalance <= 0 {
		if _, err := clients.RabbitSagaExample.BalanceRollback(&request); err != nil {
			lets.LogE("Saga: Transfer: %s", err.Error())
		}

		return &ExampleActivityTransfer{
			state: StateBalanceTranferFailed,
		}, nil
	}

	// Error response
	lets.LogI("Saga: BalanceTransfer: Failed")

	return &ExampleActivityTransfer{
		state: StateTransferFailed,
	}, err
}

// The compensation when notification fails.
type BalanceRollback struct {
	bucket *structs.SagaTransferData
}

// Execution for requesting balance rollback.
func (t *BalanceRollback) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {
	lets.LogI("Saga: BalanceRollback: Starting ...")

	var request = data.RequestBalance{
		SenderId:   t.bucket.SenderId,
		ReceiverId: t.bucket.ReceiverId,
		Amount:     t.bucket.Amount,
	}

	// Call balance rollback service provider
	if _, err := clients.RabbitSagaExample.BalanceRollback(&request); err != nil {
		lets.LogE("Saga: BalanceRollback: %s", err.Error())
	}

	// Return state where balance transaction is canceled.
	return &ExampleActivityTransfer{
		state: StateBalanceTranferCanceled,
	}, nil
}

// Notification handler.
type NotifyUser struct {
	bucket *structs.SagaTransferData
}

// Create notification in to theuser
func (t *NotifyUser) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {

	lets.LogI("Saga: NotifyUser: Starting ...")

	var request = data.RequestNotification{
		SenderId:   t.bucket.SenderId,
		ReceiverId: t.bucket.ReceiverId,
		Amount:     t.bucket.Amount,
		Status:     "success",
	}

	// Call balance rollback service provider
	if _, err := clients.RabbitSagaExample.Notification(&request); err != nil {
		lets.LogE("Saga: NotifyUser: %s", err.Error())
	}

	// Return state where noitification successfully delivered.
	return &ExampleActivityTransfer{
		state: StateNotificationCreated,
	}, nil
}

// Define saga transfer
func SagaTransfer(data *structs.SagaTransferData) (state saga.State, err error) {
	// Create saga runtime configuration.
	exampleConfig := saga.Config{
		InitState: StateTransferInit,
		Activities: []saga.Activity{
			{
				Aggregator: &ExampleTransfer{
					bucket: data,
				},
				SuccessState: StateTransferSucceed,
				FailureState: StateTransferFailed,
				Compensation: &TransferRollback{
					bucket: data,
				},
				RolledBackState: StateTransferCanceled,
			},
			{
				Aggregator: &BalanceTransfer{
					bucket: data,
				},
				SuccessState: StateBalanceTranferSucceed,
				FailureState: StateBalanceTranferFailed,
				Compensation: &BalanceRollback{
					bucket: data,
				},
				RolledBackState: StateBalanceTranferCanceled,
			},
			{
				Aggregator: &NotifyUser{
					bucket: data,
				},
				SuccessState:    StateNotificationCreated,
				FailureState:    StateNotificationFailed,
				RolledBackState: StateNotificationRetry,
			},
		},
	}

	executor, err := saga.NewExecutor(exampleConfig)
	if err != nil {
		fmt.Printf("Failed to create saga executor, err: %v.\n", err)
		return
	}

	finalTx, err := executor.Execute(context.Background(), &ExampleActivityTransfer{
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
