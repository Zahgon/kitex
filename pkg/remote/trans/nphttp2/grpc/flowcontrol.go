package grpc

import (
	"sync"
)

type writeQuota struct {
	quota int32

	ch chan struct{}

	done <-chan struct{}

	replenish func(n int)
}

func newWriteQuota(sz int32, done <-chan struct{}) *writeQuota {
	_ = "STUB: not implemented"
	return nil
}

func (w *writeQuota) get(sz int32) error { _ = "STUB: not implemented"; return nil }

func (w *writeQuota) realReplenish(n int) { _ = "STUB: not implemented"; return }

type trInFlow struct {
	limit               uint32
	unacked             uint32
	effectiveWindowSize uint32
}

func (f *trInFlow) newLimit(n uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func (f *trInFlow) onData(n uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func (f *trInFlow) reset() uint32 { _ = "STUB: not implemented"; return 0 }

func (f *trInFlow) updateEffectiveWindowSize() { _ = "STUB: not implemented"; return }

type inFlow struct {
	mu sync.Mutex

	limit uint32

	pendingData uint32

	pendingUpdate uint32

	delta uint32
}

func (f *inFlow) newLimit(n uint32) { _ = "STUB: not implemented"; return }

func (f *inFlow) maybeAdjust(n uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func (f *inFlow) onData(n uint32) error { _ = "STUB: not implemented"; return nil }

func (f *inFlow) onRead(n uint32) uint32 { _ = "STUB: not implemented"; return 0 }
