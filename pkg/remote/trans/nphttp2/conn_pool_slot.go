package nphttp2

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"
)

const (
	slotClosed int32 = 1
)

type transports struct {
	index uint32
	size  uint32

	transportSlots []*transportSlot
}

func newTransports(size uint32) *transports { _ = "STUB: not implemented"; return nil }

func (t *transports) getNextTransportSlot() (*transportSlot, uint32) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (t *transports) getTransportSlot(idx uint32) *transportSlot {
	_ = "STUB: not implemented"
	return nil
}

func (t *transports) getActiveTransport() (grpc.ClientTransport, uint32) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientTransport), 0
}

func (t *transports) createTransport(idx uint32, remoteService string,
	dialer remote.Dialer, network, address string, connectTimeout time.Duration, opts grpc.ConnectOptions,
) (grpc.ClientTransport, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientTransport), nil
}

func (t *transports) loadAll() []grpc.ClientTransport { _ = "STUB: not implemented"; return nil }

func (t *transports) close() { _ = "STUB: not implemented"; return }

var errTransportsClosed = status.Err(codes.Aborted, "transports have been closed due to instance offline")

type transportRef struct {
	ct grpc.ClientTransport
}

type transportSlot struct {
	mu    sync.Mutex
	state int32

	ref atomic.Pointer[transportRef]
}

func (slot *transportSlot) load() grpc.ClientTransport {
	_ = "STUB: not implemented"
	return *new(grpc.ClientTransport)
}

func (slot *transportSlot) createTransport(
	remoteService string,
	dialer remote.Dialer, network, address string, connectTimeout time.Duration, opts grpc.ConnectOptions,
) (grpc.ClientTransport, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientTransport), nil
}

func (slot *transportSlot) removeTransport(trans grpc.ClientTransport) {
	_ = "STUB: not implemented"
	return
}

func (slot *transportSlot) close() { _ = "STUB: not implemented"; return }
