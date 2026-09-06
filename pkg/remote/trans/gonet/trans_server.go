package gonet

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
)

const defaultShutdownTicker = 100 * time.Millisecond

func NewTransServerFactory() remote.TransServerFactory {
	_ = "STUB: not implemented"
	return *new(remote.TransServerFactory)
}

type gonetTransServerFactory struct{}

func (f *gonetTransServerFactory) NewTransServer(opt *remote.ServerOption, transHdlr remote.ServerTransHandler) remote.TransServer {
	_ = "STUB: not implemented"
	return *new(remote.TransServer)
}

type transServer struct {
	opt       *remote.ServerOption
	transHdlr remote.ServerTransHandler
	ln        net.Listener
	lncfg     net.ListenConfig
	connCount utils.AtomicInt
	shutdown  uint32
	sync.Mutex
}

var _ remote.TransServer = &transServer{}

func (ts *transServer) CreateListener(addr net.Addr) (ln net.Listener, err error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (ts *transServer) BootstrapServer(ln net.Listener) error {
	_ = "STUB: not implemented"
	return nil
}

func (ts *transServer) serveConn(ctx context.Context, conn net.Conn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ts *transServer) Shutdown() (err error) { _ = "STUB: not implemented"; return nil }

func (ts *transServer) ConnCount() utils.AtomicInt {
	_ = "STUB: not implemented"
	return *new(utils.AtomicInt)
}

func (ts *transServer) onError(ctx context.Context, err error, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (ts *transServer) refreshIdleDeadline(conn net.Conn) { _ = "STUB: not implemented"; return }

func (ts *transServer) refreshReadDeadline(ri rpcinfo.RPCInfo, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func transRecover(ctx context.Context, conn net.Conn, funcName string) {
	_ = "STUB: not implemented"
	return
}
