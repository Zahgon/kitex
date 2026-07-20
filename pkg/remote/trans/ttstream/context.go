package ttstream

import (
	"context"
	"sync/atomic"
)

type contextWithCancelReason struct {
	context.Context

	cancel context.CancelFunc
	reason atomic.Value
}

func (c *contextWithCancelReason) Err() error { _ = "STUB: not implemented"; return nil }

func (c *contextWithCancelReason) CancelWithReason(reason error) { _ = "STUB: not implemented"; return }

type cancelWithReason func(reason error)

func newContextWithCancelReason(ctx context.Context, cancel context.CancelFunc) (context.Context, cancelWithReason) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(cancelWithReason)
}
