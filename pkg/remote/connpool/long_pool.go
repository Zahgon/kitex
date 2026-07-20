package connpool

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/pkg/warmup"
)

var (
	_ net.Conn            = &longConn{}
	_ remote.LongConnPool = &LongPool{}

	sharedTickers sync.Map
)

const (
	configDumpKey = "idle_config"
)

func getSharedTicker(p *LongPool, refreshInterval time.Duration) *utils.SharedTicker {
	_ = "STUB: not implemented"
	return nil
}

type netAddr struct {
	network string
	address string
}

func (na netAddr) Network() string { _ = "STUB: not implemented"; return "" }

func (na netAddr) String() string { _ = "STUB: not implemented"; return "" }

type longConn struct {
	net.Conn
	sync.RWMutex
	deadline time.Time
	address  string
}

func (c *longConn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *longConn) RawConn() net.Conn { _ = "STUB: not implemented"; return *new(net.Conn) }

func (c *longConn) IsActive() bool { _ = "STUB: not implemented"; return false }

func (c *longConn) Expired() bool { _ = "STUB: not implemented"; return false }

type PoolDump struct {
	IdleNum       int         `json:"idle_num"`
	ConnsDeadline []time.Time `json:"conns_deadline"`
}

func newPool(minIdle, maxIdle int, maxIdleTimeout time.Duration) *pool {
	_ = "STUB: not implemented"
	return nil
}

type pool struct {
	idleList []*longConn
	mu       sync.RWMutex

	minIdle        int
	maxIdle        int
	maxIdleTimeout time.Duration
}

func (p *pool) Get() (*longConn, bool, int) { _ = "STUB: not implemented"; return nil, false, 0 }

func (p *pool) Put(o *longConn) bool { _ = "STUB: not implemented"; return false }

func (p *pool) Evict() (evicted int) { _ = "STUB: not implemented"; return 0 }

func (p *pool) Len() int { _ = "STUB: not implemented"; return 0 }

func (p *pool) Close() int { _ = "STUB: not implemented"; return 0 }

func (p *pool) Dump() PoolDump { _ = "STUB: not implemented"; return *new(PoolDump) }

func newPeer(
	serviceName string,
	addr net.Addr,
	minIdle int,
	maxIdle int,
	maxIdleTimeout time.Duration,
	globalIdle *utils.MaxCounter,
) *peer {
	_ = "STUB: not implemented"
	return nil
}

type peer struct {
	serviceName string
	addr        net.Addr
	globalIdle  *utils.MaxCounter

	pool *pool
}

func (p *peer) Get(d remote.Dialer, timeout time.Duration, reporter Reporter, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (p *peer) Put(c *longConn) error { _ = "STUB: not implemented"; return nil }

func (p *peer) Len() int { _ = "STUB: not implemented"; return 0 }

func (p *peer) Evict() { _ = "STUB: not implemented"; return }

func (p *peer) Close() { _ = "STUB: not implemented"; return }

func NewLongPool(serviceName string, idlConfig connpool.IdleConfig) *LongPool {
	_ = "STUB: not implemented"
	return nil
}

type LongPool struct {
	reporter     Reporter
	peerMap      sync.Map
	newPeer      func(net.Addr) *peer
	globalIdle   *utils.MaxCounter
	idleConfig   connpool.IdleConfig
	sharedTicker *utils.SharedTicker
	closed       int32
}

func (lp *LongPool) Get(ctx context.Context, network, address string, opt remote.ConnOption) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (lp *LongPool) Put(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (lp *LongPool) Discard(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (lp *LongPool) Clean(network, address string) { _ = "STUB: not implemented"; return }

func (lp *LongPool) Dump() interface{} { _ = "STUB: not implemented"; return nil }

func (lp *LongPool) Close() error { _ = "STUB: not implemented"; return nil }

func (lp *LongPool) EnableReporter() { _ = "STUB: not implemented"; return }

func (lp *LongPool) WarmUp(eh warmup.ErrorHandling, wuo *warmup.PoolOption, co remote.ConnOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (lp *LongPool) Evict() { _ = "STUB: not implemented"; return }

func (lp *LongPool) Tick() { _ = "STUB: not implemented"; return }

func (lp *LongPool) getPeer(addr netAddr) *peer { _ = "STUB: not implemented"; return nil }
