package client

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

type timeoutPool struct {
	size  int32
	tasks chan *timeoutTask

	maxIdle int32

	maxIdleTime time.Duration

	mu     sync.Mutex
	ticker chan struct{}
}

func newTimeoutPool(maxIdle int, maxIdleTime time.Duration) *timeoutPool {
	_ = "STUB: not implemented"
	return nil
}

func (p *timeoutPool) Size() int32 { _ = "STUB: not implemented"; return 0 }

func (p *timeoutPool) createTicker() { _ = "STUB: not implemented"; return }

func (p *timeoutPool) createWorker(t *timeoutTask) bool { _ = "STUB: not implemented"; return false }

func (p *timeoutPool) RunTask(ctx context.Context, timeout time.Duration,
	req, resp any, ep endpoint.Endpoint,
) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

var poolTask = sync.Pool{
	New: func() any {
		return &timeoutTask{}
	},
}

type timeoutTask struct {
	ctx *timeoutContext

	wg sync.WaitGroup

	req, resp any
	ep        endpoint.Endpoint

	err atomic.Value
}

func newTimeoutTask(ctx context.Context, timeout time.Duration,
	req, resp any, ep endpoint.Endpoint,
) *timeoutTask {
	_ = "STUB: not implemented"
	return nil
}

func (t *timeoutTask) recycle() { _ = "STUB: not implemented"; return }

func (t *timeoutTask) Cancel(err error) { _ = "STUB: not implemented"; return }

func (t *timeoutTask) Run() { _ = "STUB: not implemented"; return }

func (t *timeoutTask) Wait() (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *timeoutTask) waitNoTimeout() (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

type timeoutContext struct {
	context.Context

	dl time.Time
	ch chan struct{}

	mu  sync.Mutex
	err error
}

func newTimeoutContext(ctx context.Context, timeout time.Duration) *timeoutContext {
	_ = "STUB: not implemented"
	return nil
}

func (p *timeoutContext) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (p *timeoutContext) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (p *timeoutContext) Err() error { _ = "STUB: not implemented"; return nil }

func (p *timeoutContext) Cancel(err error) { _ = "STUB: not implemented"; return }
