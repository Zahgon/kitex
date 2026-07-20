package connpool

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/remote"
)

var (
	_ remote.ConnPool     = &DummyPool{}
	_ remote.LongConnPool = &DummyPool{}
	_ Reporter            = &DummyReporter{}
)

type DummyPool struct{}

func (p *DummyPool) Get(ctx context.Context, network, address string, opt remote.ConnOption) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (p *DummyPool) Put(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (p *DummyPool) Discard(conn net.Conn) error { _ = "STUB: not implemented"; return nil }

func (p *DummyPool) Clean(network, address string) { _ = "STUB: not implemented"; return }

func (p *DummyPool) Close() error { _ = "STUB: not implemented"; return nil }

type DummyReporter struct{}

func (dcm *DummyReporter) ConnSucceed(poolType ConnectionPoolType, serviceName string, addr net.Addr) {
	_ = "STUB: not implemented"
	return
}

func (dcm *DummyReporter) ConnFailed(poolType ConnectionPoolType, serviceName string, addr net.Addr) {
	_ = "STUB: not implemented"
	return
}

func (dcm *DummyReporter) ReuseSucceed(poolType ConnectionPoolType, serviceName string, addr net.Addr) {
	_ = "STUB: not implemented"
	return
}
