package connpool

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/remote"
)

var _ remote.ConnPool = &ShortPool{}

type shortConn struct {
	net.Conn
	closed bool
}

func (sc *shortConn) Close() error { _ = "STUB: not implemented"; return nil }

func (sc *shortConn) RawConn() net.Conn { _ = "STUB: not implemented"; return *new(net.Conn) }

type ShortPool struct {
	serviceName string
	reporter    Reporter
}

func NewShortPool(serviceName string) *ShortPool { _ = "STUB: not implemented"; return nil }

func (p *ShortPool) Get(ctx context.Context, network, address string, opt remote.ConnOption) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (p *ShortPool) release(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (p *ShortPool) Put(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (p *ShortPool) Discard(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (p *ShortPool) Close() error { _ = "STUB: not implemented"; return nil }

func (p *ShortPool) EnableReporter() { _ = "STUB: not implemented"; return }
