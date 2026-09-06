package rpcinfo

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cloudwego/kitex/internal"
	"github.com/cloudwego/kitex/pkg/stats"
)

var (
	_ RPCStats          = (*rpcStats)(nil)
	_ MutableRPCStats   = (*rpcStats)(nil)
	_ internal.Reusable = (*rpcStats)(nil)
	_ internal.Reusable = (*event)(nil)

	rpcStatsPool = sync.Pool{New: func() interface{} { return newRPCStats() }}
	eventPool    = sync.Pool{New: func() interface{} { return &event{} }}

	once        sync.Once
	maxEventNum int
)

type event struct {
	state  uint32
	status stats.Status
	info   string
	event  stats.Event
	time   int64
}

func (e *event) Event() stats.Event { _ = "STUB: not implemented"; return *new(stats.Event) }

func (e *event) Status() stats.Status { _ = "STUB: not implemented"; return *new(stats.Status) }

func (e *event) Info() string { _ = "STUB: not implemented"; return "" }

func (e *event) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (e *event) IsNil() bool { _ = "STUB: not implemented"; return false }

func (e *event) zero() { _ = "STUB: not implemented"; return }

func (e *event) Recycle() { _ = "STUB: not implemented"; return }

type rpcStats struct {
	level stats.Level

	eventMap []event

	sendSize     uint64
	lastSendSize uint64
	recvSize     uint64
	lastRecvSize uint64

	err      atomic.Pointer[error]
	panicErr atomic.Pointer[any]

	copied bool
}

const (
	eventUnset    uint32 = 0b0000
	eventUpdating uint32 = 0b0001
	eventRecorded uint32 = 0b0010

	eventStale uint32 = 0b0100 | eventRecorded
)

func newRPCStats() *rpcStats { _ = "STUB: not implemented"; return nil }

func NewRPCStats() RPCStats { _ = "STUB: not implemented"; return *new(RPCStats) }

func (r *rpcStats) Record(ctx context.Context, e stats.Event, status stats.Status, info string) {
	_ = "STUB: not implemented"
	return
}

func NewEvent(e stats.Event, status stats.Status, info string) Event {
	_ = "STUB: not implemented"
	return *new(Event)
}

func (r *rpcStats) SendSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (r *rpcStats) LastSendSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (r *rpcStats) RecvSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (r *rpcStats) LastRecvSize() uint64 { _ = "STUB: not implemented"; return 0 }

func (r *rpcStats) Error() error { _ = "STUB: not implemented"; return nil }

func (r *rpcStats) Panicked() (bool, any) { _ = "STUB: not implemented"; return false, *new(any) }

func (r *rpcStats) GetEvent(e stats.Event) Event { _ = "STUB: not implemented"; return *new(Event) }

func (r *rpcStats) Level() stats.Level { _ = "STUB: not implemented"; return *new(stats.Level) }

func (r *rpcStats) CopyForRetry() RPCStats { _ = "STUB: not implemented"; return *new(RPCStats) }

func (r *rpcStats) SetSendSize(size uint64) { _ = "STUB: not implemented"; return }

func (r *rpcStats) IncrSendSize(size uint64) { _ = "STUB: not implemented"; return }

func (r *rpcStats) SetRecvSize(size uint64) { _ = "STUB: not implemented"; return }

func (r *rpcStats) IncrRecvSize(size uint64) { _ = "STUB: not implemented"; return }

func (r *rpcStats) SetError(err error) { _ = "STUB: not implemented"; return }

func (r *rpcStats) SetPanicked(x any) { _ = "STUB: not implemented"; return }

func (r *rpcStats) SetLevel(level stats.Level) { _ = "STUB: not implemented"; return }

func (r *rpcStats) Reset() { _ = "STUB: not implemented"; return }

func (r *rpcStats) ImmutableView() RPCStats { _ = "STUB: not implemented"; return *new(RPCStats) }

func (r *rpcStats) Recycle() { _ = "STUB: not implemented"; return }
