package nphttp2

import (
	"context"
	"net"

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
	opt   *remote.ClientOption
	codec remote.Codec
}

func (h *cliTransHandler) Write(ctx context.Context, conn net.Conn, msg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *cliTransHandler) Read(ctx context.Context, conn net.Conn, msg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *cliTransHandler) OnRead(ctx context.Context, conn net.Conn) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *cliTransHandler) OnConnect(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *cliTransHandler) OnInactive(ctx context.Context, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (h *cliTransHandler) OnError(ctx context.Context, err error, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (h *cliTransHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *cliTransHandler) SetPipeline(pipeline *remote.TransPipeline) {
	_ = "STUB: not implemented"
	return
}
