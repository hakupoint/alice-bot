package reptile

import "context"

type Tasks struct {
	Total int
	Run   chan func(int) error
	Error chan error
	ctx   context.Context
}

func (t *Tasks) run() {

}
