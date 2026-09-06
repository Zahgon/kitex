package netpollmux

import (
	"context"
	"net"
	"sync"

	"github.com/cloudwego/kitex/pkg/remote"

	"golang.org/x/sync/singleflight"
)

var _ remote.LongConnPool = &MuxPool{}

func NewMuxConnPool(size int) *MuxPool { _ = "STUB: not implemented"; return nil }

type MuxPool struct {
	size    int32
	sfg     singleflight.Group
	connMap sync.Map
}

func (mp *MuxPool) Get(ctx context.Context, network, address string, opt remote.ConnOption) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (mp *MuxPool) Put(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (mp *MuxPool) Discard(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (mp *MuxPool) Clean(network, address string) { _ = "STUB: not implemented"; return }

func (mp *MuxPool) Close() error { _ = "STUB: not implemented"; return nil }

type conns struct {
	index uint32
	size  uint32
	conns []*muxCliConn
}

func (c *conns) get() *muxCliConn { _ = "STUB: not implemented"; return nil }

func (c *conns) put(conn *muxCliConn) { _ = "STUB: not implemented"; return }

func (c *conns) close() { _ = "STUB: not implemented"; return }
