package manager

import (
	"github.com/bongnv/saga"
)

type ProcessTransfer struct {
	balance float64
}

func (t *ProcessTransfer) Execute(amount float64) (saga.Transaction, error) {
	t.balance -= amount

	if t.balance < 0 {
		return &Transfer{
			state: TransferFailed,
			// }, errors.New("insuficient balance")
		}, nil
	}

	return &Transfer{
		state: TransferSuccess,
	}, nil

}
