package netpoll

import (
	"context"
	"net"
	"sync"

	"github.com/cloudwego/netpoll"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/utils"
)

func NewTransServerFactory() remote.TransServerFactory {
	_ = "STUB: not implemented"
	return *new(remote.TransServerFactory)
}

type netpollTransServerFactory struct{}

func (f *netpollTransServerFactory) NewTransServer(opt *remote.ServerOption, transHdlr remote.ServerTransHandler) remote.TransServer {
	_ = "STUB: not implemented"
	return *new(remote.TransServer)
}

type transServer struct {
	opt       *remote.ServerOption
	transHdlr remote.ServerTransHandler

	evl       netpoll.EventLoop
	ln        net.Listener
	lncfg     net.ListenConfig
	connCount utils.AtomicInt
	sync.Mutex
}

var _ remote.TransServer = &transServer{}

func (ts *transServer) CreateListener(addr net.Addr) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (ts *transServer) BootstrapServer(ln net.Listener) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ts *transServer) Shutdown() (err error) { _ = "STUB: not implemented"; return nil }

func (ts *transServer) ConnCount() utils.AtomicInt {
	_ = "STUB: not implemented"
	return *new(utils.AtomicInt)
}

func (ts *transServer) onConnActive(conn netpoll.Connection) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (ts *transServer) onConnRead(ctx context.Context, conn netpoll.Connection) error {
	_ = "STUB: not implemented"
	return nil
}

func (ts *transServer) onConnInactive(ctx context.Context, conn netpoll.Connection) {
	_ = "STUB: not implemented"
	return
}

func (ts *transServer) onError(ctx context.Context, err error, conn netpoll.Connection) {
	_ = "STUB: not implemented"
	return
}

func transRecover(ctx context.Context, conn netpoll.Connection, funcName string, propagatePanic bool) {
	_ = "STUB: not implemented"
	return
}
