package detection

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/remote"
)

type DetectableServerTransHandler interface {
	remote.ServerTransHandler
	ProtocolMatch(ctx context.Context, conn net.Conn) (err error)
}

func NewSvrTransHandlerFactory(defaultHandlerFactory remote.ServerTransHandlerFactory,
	detectableHandlerFactory ...remote.ServerTransHandlerFactory,
) remote.ServerTransHandlerFactory {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandlerFactory)
}

type svrTransHandlerFactory struct {
	defaultHandlerFactory    remote.ServerTransHandlerFactory
	detectableHandlerFactory []remote.ServerTransHandlerFactory
}

func (f *svrTransHandlerFactory) NewTransHandler(opt *remote.ServerOption) (remote.ServerTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandler), nil
}

type svrTransHandler struct {
	defaultHandler remote.ServerTransHandler
	registered     []DetectableServerTransHandler
}

func (t *svrTransHandler) Write(ctx context.Context, conn net.Conn, send remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) Read(ctx context.Context, conn net.Conn, msg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) OnRead(ctx context.Context, conn net.Conn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (t *svrTransHandler) OnInactive(ctx context.Context, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) OnError(ctx context.Context, err error, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) which(ctx context.Context) remote.ServerTransHandler {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandler)
}

func (t *svrTransHandler) SetPipeline(pipeline *remote.TransPipeline) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) SetInvokeHandleFunc(inkHdlFunc endpoint.Endpoint) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) GracefulShutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type handlerKey struct{}

type handlerWrapper struct {
	ctx     context.Context
	handler remote.ServerTransHandler
}
