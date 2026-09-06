package limiter

import (
	"context"
	"time"
)

type DummyConcurrencyLimiter struct{}

func (dcl *DummyConcurrencyLimiter) Acquire(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (dcl *DummyConcurrencyLimiter) Release(ctx context.Context) { _ = "STUB: not implemented"; return }

func (dcl *DummyConcurrencyLimiter) Status(ctx context.Context) (limit, occupied int) {
	_ = "STUB: not implemented"
	return 0, 0
}

type DummyRateLimiter struct{}

func (drl *DummyRateLimiter) Acquire(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (drl *DummyRateLimiter) Status(ctx context.Context) (max, current int, interval time.Duration) {
	_ = "STUB: not implemented"
	return 0, 0, *new(time.Duration)
}
