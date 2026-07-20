package ttstream

import (
	"context"
	"net"
	"sync"

	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/netpoll"

	"github.com/cloudwego/kitex/pkg/remote/trans/ttstream/internal/container"
)

type serverTransport struct {
	conn netpoll.Connection

	streams sync.Map
	scache  []*serverStream
	spipe   *container.Pipe[*serverStream]

	mu        sync.Mutex
	state     int32
	closedErr error
	writer    *writerBuffer

	closedTrigger chan struct{}
}

func newServerTransport(conn netpoll.Connection) *serverTransport {
	_ = "STUB: not implemented"
	return nil
}

func (t *serverTransport) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (t *serverTransport) Close(exception error) error { _ = "STUB: not implemented"; return nil }

func (t *serverTransport) setClosedStateLocked(err error) { _ = "STUB: not implemented"; return }

func (t *serverTransport) releaseResources(err error) { _ = "STUB: not implemented"; return }

func (t *serverTransport) WaitClosed() { _ = "STUB: not implemented"; return }

func (t *serverTransport) IsActive() bool { _ = "STUB: not implemented"; return false }

func (t *serverTransport) storeStream(s *serverStream) { _ = "STUB: not implemented"; return }

func (t *serverTransport) loadStream(sid int32) (s *serverStream, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (t *serverTransport) deleteStream(sid int32) { _ = "STUB: not implemented"; return }

func (t *serverTransport) readFrame(reader bufiox.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *serverTransport) loopRead() error { _ = "STUB: not implemented"; return nil }

func (t *serverTransport) WriteFrame(fr *Frame) (err error) { _ = "STUB: not implemented"; return nil }

func (t *serverTransport) CloseStream(sid int32) (err error) { _ = "STUB: not implemented"; return nil }

func (t *serverTransport) ReadStream(ctx context.Context) (*serverStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
