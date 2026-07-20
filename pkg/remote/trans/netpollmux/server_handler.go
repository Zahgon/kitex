package netpollmux

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/cloudwego/netpoll"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

const defaultExitWaitGracefulShutdownTime = 1 * time.Second

type svrTransHandlerFactory struct{}

func NewSvrTransHandlerFactory() remote.ServerTransHandlerFactory {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandlerFactory)
}

func (f *svrTransHandlerFactory) MuxEnabled() bool { _ = "STUB: not implemented"; return false }

func (f *svrTransHandlerFactory) NewTransHandler(opt *remote.ServerOption) (remote.ServerTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandler), nil
}

func newSvrTransHandler(opt *remote.ServerOption) (*svrTransHandler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ remote.ServerTransHandler = &svrTransHandler{}

type svrTransHandler struct {
	opt         *remote.ServerOption
	svcSearcher remote.ServiceSearcher
	inkHdlFunc  endpoint.Endpoint
	codec       remote.Codec
	transPipe   *remote.TransPipeline
	ext         trans.Extension
	funcPool    sync.Pool
	conns       sync.Map
	tasks       sync.WaitGroup
}

func (t *svrTransHandler) Write(ctx context.Context, conn net.Conn, sendMsg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) Read(ctx context.Context, conn net.Conn, msg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) readWithByteBuffer(ctx context.Context, bufReader remote.ByteBuffer, msg remote.Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (t *svrTransHandler) OnRead(muxSvrConnCtx context.Context, conn net.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *svrTransHandler) batchGoTasks(fs []func()) { _ = "STUB: not implemented"; return }

func (t *svrTransHandler) task(muxSvrConnCtx context.Context, conn net.Conn, reader netpoll.Reader) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

type ctxKeyMuxSvrConn struct{}

func (t *svrTransHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) OnInactive(ctx context.Context, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) GracefulShutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *svrTransHandler) OnError(ctx context.Context, err error, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) SetInvokeHandleFunc(inkHdlFunc endpoint.Endpoint) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) SetPipeline(p *remote.TransPipeline) { _ = "STUB: not implemented"; return }

func (t *svrTransHandler) writeErrorReplyIfNeeded(
	ctx context.Context, recvMsg remote.Message, conn net.Conn, ri rpcinfo.RPCInfo, err error, doOnMessage bool,
) (shouldCloseConn bool) {
	_ = "STUB: not implemented"
	return false
}

func (t *svrTransHandler) tryRecover(ctx context.Context, conn net.Conn) {
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

func getRemoteInfo(ri rpcinfo.RPCInfo, conn net.Conn) (string, net.Addr) {
	_ = "STUB: not implemented"
	return "", *new(net.Addr)
}
