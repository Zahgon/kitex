package utils

import (
	"sync"
	"time"
)

type TickerTask interface {
	Tick()
}

func NewSharedTicker(interval time.Duration) *SharedTicker { _ = "STUB: not implemented"; return nil }

func NewSyncSharedTicker(interval time.Duration) *SharedTicker {
	_ = "STUB: not implemented"
	return nil
}

type SharedTicker struct {
	sync.Mutex
	started  bool
	Interval time.Duration
	tasks    map[TickerTask]struct{}
	stopChan chan struct{}
	sync     bool
}

func (t *SharedTicker) Add(b TickerTask) { _ = "STUB: not implemented"; return }

func (t *SharedTicker) Delete(b TickerTask) { _ = "STUB: not implemented"; return }

func (t *SharedTicker) Closed() bool { _ = "STUB: not implemented"; return false }

func (t *SharedTicker) Tick(interval time.Duration) { _ = "STUB: not implemented"; return }

func (t *SharedTicker) syncExec() { _ = "STUB: not implemented"; return }

func (t *SharedTicker) asyncExec() { _ = "STUB: not implemented"; return }
