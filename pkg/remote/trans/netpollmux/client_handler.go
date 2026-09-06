package netpollmux

import (
	"context"
	"net"

	"github.com/cloudwego/netpoll"

	"github.com/cloudwego/kitex/pkg/remote"
)

type cliTransHandlerFactory struct{}

func NewCliTransHandlerFactory() remote.ClientTransHandlerFactory {
	_ = "STUB: not implemented"
	return *new(remote.ClientTransHandlerFactory)
}

func (f *cliTransHandlerFactory) NewTransHandler(opt *remote.ClientOption) (remote.ClientTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ClientTransHandler), nil
}

func newCliTransHandler(opt *remote.ClientOption) (*cliTransHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ remote.ClientTransHandler = &cliTransHandler{}

type cliTransHandler struct {
	opt       *remote.ClientOption
	codec     remote.Codec
	transPipe *remote.TransPipeline
}

func (t *cliTransHandler) Write(ctx context.Context, conn net.Conn, sendMsg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *cliTransHandler) Read(ctx context.Context, conn net.Conn, msg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *cliTransHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *cliTransHandler) OnInactive(ctx context.Context, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (t *cliTransHandler) OnError(ctx context.Context, err error, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (t *cliTransHandler) SetPipeline(p *remote.TransPipeline) { _ = "STUB: not implemented"; return }

type asyncCallback struct {
	wbuf       *netpoll.LinkBuffer
	bufWriter  remote.ByteBuffer
	notifyChan chan remote.ByteBuffer
	closed     int32
}

func newAsyncCallback(wbuf *netpoll.LinkBuffer, bufWriter remote.ByteBuffer) *asyncCallback {
	_ = "STUB: not implemented"
	return nil
}

func (c *asyncCallback) Recv(bufReader remote.ByteBuffer, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *asyncCallback) Close() error { _ = "STUB: not implemented"; return nil }

func (c *asyncCallback) notify(bufReader remote.ByteBuffer) { _ = "STUB: not implemented"; return }

func (c *asyncCallback) getter() (w netpoll.Writer, isNil bool) {
	_ = "STUB: not implemented"
	return *new(netpoll.Writer), false
}
