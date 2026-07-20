package remote

import (
	"net"

	"github.com/cloudwego/kitex/pkg/utils"
)

type TransServerFactory interface {
	NewTransServer(opt *ServerOption, transHdlr ServerTransHandler) TransServer
}

type TransServer interface {
	CreateListener(net.Addr) (net.Listener, error)
	BootstrapServer(net.Listener) (err error)
	Shutdown() error
	ConnCount() utils.AtomicInt
}
