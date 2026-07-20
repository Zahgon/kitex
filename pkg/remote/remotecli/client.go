package remotecli

import (
	"context"
	"net"
	"sync"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type Client interface {
	Send(ctx context.Context, ri rpcinfo.RPCInfo, req remote.Message) (err error)
	Recv(ctx context.Context, ri rpcinfo.RPCInfo, resp remote.Message) (err error)
	Recycle()
}

var clientPool = &sync.Pool{
	New: func() interface{} {
		return new(client)
	},
}

type client struct {
	transHdlr   remote.TransHandler
	connManager *ConnWrapper
	conn        net.Conn
}

func NewClient(ctx context.Context, ri rpcinfo.RPCInfo, handler remote.TransHandler, opt *remote.ClientOption) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func (c *client) Recycle() { _ = "STUB: not implemented"; return }

func (c *client) init(handler remote.TransHandler, cm *ConnWrapper, conn net.Conn) {
	c.transHdlr = handler
	c.connManager = cm
	c.conn = conn
}

func (c *client) Send(ctx context.Context, ri rpcinfo.RPCInfo, req remote.Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) Recv(ctx context.Context, ri rpcinfo.RPCInfo, resp remote.Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}
