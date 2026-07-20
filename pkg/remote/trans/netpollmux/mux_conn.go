package netpollmux

import (
	"context"
	"errors"
	"net"
	"sync"

	"github.com/cloudwego/netpoll"
	"github.com/cloudwego/netpoll/mux"

	"github.com/cloudwego/kitex/pkg/remote/codec"
)

var ErrConnClosed = errors.New("conn closed")

var defaultCodec = codec.NewDefaultCodec()

func newMuxCliConn(connection netpoll.Connection) *muxCliConn {
	_ = "STUB: not implemented"
	return nil
}

type muxCliConn struct {
	muxConn
	closing  bool
	seqIDMap *shardMap
}

func (c *muxCliConn) IsActive() bool { _ = "STUB: not implemented"; return false }

func (c *muxCliConn) OnRequest(ctx context.Context, connection netpoll.Connection) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *muxCliConn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *muxCliConn) forceClose() error { _ = "STUB: not implemented"; return nil }

func (c *muxCliConn) close() error { _ = "STUB: not implemented"; return nil }

func (c *muxCliConn) onError(ctx context.Context, err error, connection netpoll.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

func newMuxSvrConn(connection netpoll.Connection, pool *sync.Pool) *muxSvrConn {
	_ = "STUB: not implemented"
	return nil
}

type muxSvrConn struct {
	muxConn
	pool *sync.Pool
}

func newMuxConn(connection netpoll.Connection) muxConn {
	_ = "STUB: not implemented"
	return *new(muxConn)
}

var (
	_ net.Conn           = &muxConn{}
	_ netpoll.Connection = &muxConn{}
)

type muxConn struct {
	netpoll.Connection
	shardQueue *mux.ShardQueue
}

func (c *muxConn) Put(gt mux.WriterGetter) { _ = "STUB: not implemented"; return }

func (c *muxConn) GracefulShutdown() { _ = "STUB: not implemented"; return }
