package limiter

import (
	"context"
	"time"
)

var fixedWindowTime = time.Second

type qpsLimiter struct {
	limit      int32
	tokens     int32
	interval   time.Duration
	once       int32
	ticker     *time.Ticker
	tickerDone chan bool
}

func NewQPSLimiter(interval time.Duration, limit int) RateLimiter {
	_ = "STUB: not implemented"
	return *new(RateLimiter)
}

func (l *qpsLimiter) UpdateLimit(limit int) { _ = "STUB: not implemented"; return }

func (l *qpsLimiter) UpdateQPSLimit(interval time.Duration, limit int) {
	_ = "STUB: not implemented"
	return
}

func (l *qpsLimiter) Acquire(ctx context.Context) bool { _ = "STUB: not implemented"; return false }

func (l *qpsLimiter) Status(ctx context.Context) (max, cur int, interval time.Duration) {
	_ = "STUB: not implemented"
	return 0, 0, *new(time.Duration)
}

func (l *qpsLimiter) startTicker(interval time.Duration) { _ = "STUB: not implemented"; return }

func (l *qpsLimiter) stopTicker() { _ = "STUB: not implemented"; return }

func (l *qpsLimiter) updateToken() { _ = "STUB: not implemented"; return }

func calcOnce(interval time.Duration, limit int) int32 { _ = "STUB: not implemented"; return 0 }

func (l *qpsLimiter) resetTokens(once int32) { _ = "STUB: not implemented"; return }
