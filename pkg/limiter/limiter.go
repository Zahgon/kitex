package limiter

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/limit"
)

type ConcurrencyLimiter interface {
	Acquire(ctx context.Context) bool

	Release(ctx context.Context)

	Status(ctx context.Context) (limit, occupied int)
}

type RateLimiter interface {
	Acquire(ctx context.Context) bool

	Status(ctx context.Context) (max, current int, interval time.Duration)
}

type Updatable interface {
	UpdateLimit(limit int)
}

type LimitReporter interface {
	ConnOverloadReport()
	QPSOverloadReport()
}

func NewLimiterWrapper(conLimit ConcurrencyLimiter, qpsLimit RateLimiter) limit.Updater {
	_ = "STUB: not implemented"
	return *new(limit.Updater)
}

type limitWrapper struct {
	conLimit ConcurrencyLimiter
	qpsLimit RateLimiter
}

func (l *limitWrapper) UpdateLimit(opt *limit.Option) (updated bool) {
	_ = "STUB: not implemented"
	return false
}
