package warmup

import (
	"context"
	"sync"

	"github.com/cloudwego/kitex/pkg/remote"
)

type PoolHelper struct {
	ErrorHandling
}

func (p *PoolHelper) WarmUp(po *PoolOption, pool remote.ConnPool, co remote.ConnOption) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type manager struct {
	ErrorHandling
	work    context.Context
	control context.Context
	cancel  context.CancelFunc
	errs    chan *job
	jobs    chan *job
	bads    []*job
	errLock sync.RWMutex
}

func newManager(ctx context.Context, po *PoolOption, eh ErrorHandling) *manager {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) report() error { _ = "STUB: not implemented"; return nil }

func (m *manager) watch() { _ = "STUB: not implemented"; return }

func discard(ch chan *job) { _ = "STUB: not implemented"; return }

func split(targets map[string][]string, dup int) (js chan *job) {
	_ = "STUB: not implemented"
	return nil
}

type job struct {
	network string
	address string
	err     error
}

type worker struct {
	jobs chan *job
	errs chan *job
	pool remote.ConnPool
	co   remote.ConnOption
	po   *PoolOption
}

func (w *worker) fire(ctx context.Context) { _ = "STUB: not implemented"; return }
