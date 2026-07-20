package remotesvr

import (
	"net"
	"sync"

	"github.com/cloudwego/kitex/pkg/remote"
)

type Server interface {
	Start() chan error
	Stop() error
	Address() net.Addr
}

type server struct {
	opt      *remote.ServerOption
	listener net.Listener
	transSvr remote.TransServer
	sync.Mutex
}

func NewServer(opt *remote.ServerOption, transHdlr remote.ServerTransHandler) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

func (s *server) Start() chan error { _ = "STUB: not implemented"; return nil }

func (s *server) buildListener() (ln net.Listener, err error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func (s *server) Stop() (err error) { _ = "STUB: not implemented"; return nil }

func (s *server) Address() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }
