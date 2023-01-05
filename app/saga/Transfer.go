package saga

import (
	"context"
	"fmt"

	"github.com/bongnv/saga"
)

const (
	stateTransferInit saga.State = iota
	stateTransferSucceed
	stateTransferFailed
	stateTransferCanceled
	stateNotifyUserInit
	stateNotifyUserSucceed
	stateNotifyUserFailed
	stateNotifyUserRetry
)

type ActivityTransfer struct {
	state saga.State
}

func (t *ActivityTransfer) State() saga.State {
	return t.state
}

type TransferStart struct {
	from   string
	to     string
	amount string
}

func (a *TransferStart) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {

	// Publish to

	return &ActivityTransfer{
		state: stateTransferSucceed,
	}, nil
}

type TransferCancel struct {
	from   string
	to     string
	amount string
}

func (a *TransferCancel) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {

	// Publish to

	return &ActivityTransfer{
		state: stateTransferCanceled,
	}, nil
}

type NotifyUserStart struct{}

func (a *NotifyUserStart) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {

	// Publish to

	return &ActivityTransfer{
		state: stateTransferSucceed,
	}, nil
}

type NotifyUserRetry struct{}

func (a *NotifyUserRetry) Execute(ctx context.Context, tx saga.Transaction) (saga.Transaction, error) {

	// Publish to

	return &ActivityTransfer{
		state: stateTransferCanceled,
	}, nil
}

// Define saga transfer
func SagaTransfer() {
	exampleConfig := saga.Config{
		InitState: stateTransferInit,
		Activities: []saga.Activity{
			{
				Aggregator:      &TransferStart{},
				SuccessState:    stateTransferSucceed,
				FailureState:    stateTransferFailed,
				Compensation:    &TransferCancel{},
				RolledBackState: stateTransferCanceled,
			},
			{
				Aggregator:      &NotifyUserStart{},
				SuccessState:    stateNotifyUserInit,
				FailureState:    stateNotifyUserFailed,
				RolledBackState: stateNotifyUserRetry,
			},
		},
	}

	executor, err := saga.NewExecutor(exampleConfig)
	if err != nil {
		fmt.Printf("Failed to create saga executor, err: %v.\n", err)
		return
	}

	finalTx, err := executor.Execute(context.Background(), &ActivityTransfer{
		state: stateTransferInit,
	})
	if err != nil {
		fmt.Printf("Failed to execute saga transaction, err: %v.\n", err)
		return
	}

	if finalTx.State() == stateTransferCanceled {
		fmt.Println("OK")
	}
}
