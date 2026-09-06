package nphttp2

import (
	"context"
	"crypto/tls"
	"net"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"
)

const (
	poolOpen   int32 = 0
	poolClosed int32 = 1
)

func poolSize() uint32 { _ = "STUB: not implemented"; return 0 }

func NewConnPool(remoteService string, size uint32, connOpts grpc.ConnectOptions) *connPool {
	_ = "STUB: not implemented"
	return nil
}

type connPool struct {
	size          uint32
	sfg           singleflight.Group
	conns         sync.Map
	remoteService string
	connOpts      grpc.ConnectOptions
	closed        int32
}

var (
	_                 remote.LongConnPool = (*connPool)(nil)
	errConnPoolClosed                     = status.Err(codes.Aborted, "connection pool has been closed")
)

func (p *connPool) Get(ctx context.Context, network, address string, opt remote.ConnOption) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (p *connPool) Put(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (p *connPool) release(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (p *connPool) createShortConn(ctx context.Context, network, address string, opt remote.ConnOption) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (p *connPool) Discard(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (p *connPool) Clean(network, address string) { _ = "STUB: not implemented"; return }

func (p *connPool) Close() error { _ = "STUB: not implemented"; return nil }

func (p *connPool) isClosed() bool { _ = "STUB: not implemented"; return false }

func (p *connPool) casClosed() bool { _ = "STUB: not implemented"; return false }

type dumpEntry struct {
	addr string
	tr   grpc.ClientTransport
}

func (p *connPool) Dump() interface{} { _ = "STUB: not implemented"; return nil }

func newTransport(remoteService string,
	dialer remote.Dialer, network, address string, connectTimeout time.Duration, opts grpc.ConnectOptions,
	onGoAway func(context.Context, grpc.ClientTransport, grpc.GoAwayReason),
	onClose func(context.Context, grpc.ClientTransport, error),
) (grpc.ClientTransport, error) {
	_ = "STUB: not implemented"
	return *new(grpc.ClientTransport), nil
}

func newTLSConn(conn net.Conn, tlsCfg *tls.Config) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func checkActive(trans grpc.ClientTransport) bool { _ = "STUB: not implemented"; return false }
