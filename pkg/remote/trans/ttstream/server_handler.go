package ttstream

import (
	"context"
	"errors"
	"net"

	"github.com/cloudwego/netpoll"

	igeneric "github.com/cloudwego/kitex/internal/generic"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

var streamingBidirectionalCtx = igeneric.WithGenericStreamingMode(context.Background(), serviceinfo.StreamingBidirectional)

type (
	serverTransCtxKey struct{}
)

type svrTransHandlerFactory struct{}

func NewSvrTransHandlerFactory() remote.ServerTransHandlerFactory {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandlerFactory)
}

func (f *svrTransHandlerFactory) NewTransHandler(opts *remote.ServerOption) (remote.ServerTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandler), nil
}

var (
	_                   remote.ServerTransHandler = &svrTransHandler{}
	errProtocolNotMatch                           = errors.New("protocol not match")
	errNilTransport                               = errors.New("server transport is nil")
)

type svrTransHandler struct {
	opt           *remote.ServerOption
	inkHdlFunc    endpoint.Endpoint
	headerHandler HeaderFrameReadHandler
}

func (t *svrTransHandler) SetInvokeHandleFunc(inkHdlFunc endpoint.Endpoint) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) ProtocolMatch(ctx context.Context, conn net.Conn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type onDisConnectSetter interface {
	SetOnDisconnect(onDisconnect netpoll.OnDisconnect) error
}

func (t *svrTransHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) OnRead(ctx context.Context, conn net.Conn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (t *svrTransHandler) OnStream(ctx context.Context, conn net.Conn, st *serverStream) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:staticcheck // SA1029: consts.CtxKeyMethod has been used and we just follow it

//nolint:staticcheck // SA1019: intentional internal rollout helper usage before single-hop becomes default.

func (t *svrTransHandler) Write(ctx context.Context, conn net.Conn, send remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) Read(ctx context.Context, conn net.Conn, msg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) OnInactive(ctx context.Context, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) OnError(ctx context.Context, err error, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) OnStreamFinish(ss streaming.ServerStream, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *svrTransHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) SetPipeline(pipeline *remote.TransPipeline) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) startTracer(ctx context.Context, ri rpcinfo.RPCInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (t *svrTransHandler) finishTracer(ctx context.Context, ri rpcinfo.RPCInfo, err error, panicErr interface{}) {
	_ = "STUB: not implemented"
	return
}
