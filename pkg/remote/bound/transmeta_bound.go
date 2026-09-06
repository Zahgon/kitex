package bound

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/remote"
)

func NewTransMetaHandler(mhs []remote.MetaHandler) remote.DuplexBoundHandler {
	_ = "STUB: not implemented"
	return *new(remote.DuplexBoundHandler)
}

type transMetaHandler struct {
	mhs []remote.MetaHandler
}

func (h *transMetaHandler) Write(ctx context.Context, conn net.Conn, sendMsg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *transMetaHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *transMetaHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *transMetaHandler) OnRead(ctx context.Context, conn net.Conn) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (h *transMetaHandler) OnInactive(ctx context.Context, conn net.Conn) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getValidMsg(args, result remote.Message) (msg remote.Message, isServer bool) {
	_ = "STUB: not implemented"
	return *new(remote.Message), false
}
