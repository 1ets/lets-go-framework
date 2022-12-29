package manager

import "github.com/bongnv/saga"

const (
	TransferInit saga.State = iota
	TransferProcess
	TransferSuccess
	TransferFailed
)

type Transfer struct {
	state saga.State
}

func (t *Transfer) State() saga.State {
	return t.state
}
