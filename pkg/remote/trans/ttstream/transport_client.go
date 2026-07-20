package ttstream

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/netpoll"

	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/pkg/utils"
)

var ticker = utils.NewSyncSharedTicker(5 * time.Second)

type clientTransport struct {
	conn    netpoll.Connection
	pool    transPool
	streams sync.Map

	mu        sync.Mutex
	state     int32
	closedErr error
	writer    *writerBuffer

	closedTrigger chan struct{}
}

func newClientTransport(conn netpoll.Connection, pool transPool) *clientTransport {
	_ = "STUB: not implemented"
	return nil
}

func (t *clientTransport) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (t *clientTransport) Close(exception error) error { _ = "STUB: not implemented"; return nil }

func (t *clientTransport) setClosedStateLocked(err error) { _ = "STUB: not implemented"; return }

func (t *clientTransport) releaseResources(err error) { _ = "STUB: not implemented"; return }

func (t *clientTransport) WaitClosed() { _ = "STUB: not implemented"; return }

func (t *clientTransport) IsActive() bool { _ = "STUB: not implemented"; return false }

func (t *clientTransport) storeStream(s *clientStream) { _ = "STUB: not implemented"; return }

func (t *clientTransport) loadStream(sid int32) (s *clientStream, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (t *clientTransport) deleteStream(sid int32) { _ = "STUB: not implemented"; return }

func (t *clientTransport) readFrame(reader bufiox.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *clientTransport) loopRead() error { _ = "STUB: not implemented"; return nil }

func (t *clientTransport) WriteFrame(fr *Frame) (err error) { _ = "STUB: not implemented"; return nil }

func (t *clientTransport) CloseStream(sid int32) (err error) { _ = "STUB: not implemented"; return nil }

var clientStreamID int32

func genStreamID() int32 { _ = "STUB: not implemented"; return 0 }

func (t *clientTransport) WriteStream(
	ctx context.Context, s *clientStream, intHeader IntHeader, strHeader streaming.Header,
) error {
	_ = "STUB: not implemented"
	return nil
}
