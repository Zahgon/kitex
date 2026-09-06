package limiter

import (
	"context"
)

type connectionLimiter struct {
	lim  int32
	curr int32
}

func NewConnectionLimiter(lim int) ConcurrencyLimiter {
	_ = "STUB: not implemented"
	return *new(ConcurrencyLimiter)
}

func NewConcurrencyLimiter(lim int) ConcurrencyLimiter {
	_ = "STUB: not implemented"
	return *new(ConcurrencyLimiter)
}

func (ml *connectionLimiter) Acquire(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (ml *connectionLimiter) Release(ctx context.Context) { _ = "STUB: not implemented"; return }

func (ml *connectionLimiter) UpdateLimit(lim int) { _ = "STUB: not implemented"; return }

func (ml *connectionLimiter) Status(ctx context.Context) (limit, occupied int) {
	_ = "STUB: not implemented"
	return 0, 0
}
