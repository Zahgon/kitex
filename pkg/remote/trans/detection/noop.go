package detection

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/remote"
)

var noopHandler = noopSvrTransHandler{}

type noopSvrTransHandler struct{}

func (noopSvrTransHandler) Write(ctx context.Context, conn net.Conn, send remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (noopSvrTransHandler) Read(ctx context.Context, conn net.Conn, msg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (noopSvrTransHandler) OnRead(ctx context.Context, conn net.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (noopSvrTransHandler) OnInactive(ctx context.Context, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}
func (noopSvrTransHandler) OnError(ctx context.Context, err error, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (noopSvrTransHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (noopSvrTransHandler) SetPipeline(pipeline *remote.TransPipeline) {
	_ = "STUB: not implemented"
	return
}
func (noopSvrTransHandler) SetInvokeHandleFunc(inkHdlFunc endpoint.Endpoint) {
	_ = "STUB: not implemented"
	return
}
func (noopSvrTransHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
