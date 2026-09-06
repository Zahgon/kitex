package remotecli

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var connWrapperPool sync.Pool

func init() {
	connWrapperPool.New = newConnWrapper
}

var _ ConnReleaser = &ConnWrapper{}

type ConnReleaser interface {
	ReleaseConn(err error, ri rpcinfo.RPCInfo)
}

type ConnWrapper struct {
	connPool remote.ConnPool
	conn     net.Conn
}

func NewConnWrapper(connPool remote.ConnPool) *ConnWrapper { _ = "STUB: not implemented"; return nil }

func (cm *ConnWrapper) GetConn(ctx context.Context, d remote.Dialer, ri rpcinfo.RPCInfo) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (cm *ConnWrapper) ReleaseConn(err error, ri rpcinfo.RPCInfo) {
	_ = "STUB: not implemented"
	return
}

func newConnWrapper() interface{} { _ = "STUB: not implemented"; return nil }

func (cm *ConnWrapper) zero() { _ = "STUB: not implemented"; return }

func (cm *ConnWrapper) getConnWithPool(ctx context.Context, cp remote.ConnPool, d remote.Dialer,
	timeout time.Duration, ri rpcinfo.RPCInfo,
) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (cm *ConnWrapper) getConnWithDialer(ctx context.Context, d remote.Dialer,
	timeout time.Duration, ri rpcinfo.RPCInfo,
) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
