package trans

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/internal/logbackoff"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func NewDefaultSvrTransHandler(opt *remote.ServerOption, ext Extension) (remote.ServerTransHandler, error) {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandler), nil
}

type svrTransHandler struct {
	opt                *remote.ServerOption
	svcSearcher        remote.ServiceSearcher
	inkHdlFunc         endpoint.Endpoint
	codec              remote.Codec
	transPipe          *remote.TransPipeline
	ext                Extension
	inGracefulShutdown uint32
	remoteClosedWarn   logbackoff.Exponential
}

func (t *svrTransHandler) Write(ctx context.Context, conn net.Conn, sendMsg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) Read(ctx context.Context, conn net.Conn, recvMsg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) newCtxWithRPCInfo(ctx context.Context, conn net.Conn) (context.Context, rpcinfo.RPCInfo) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(rpcinfo.RPCInfo)
}

func (t *svrTransHandler) OnRead(ctx context.Context, conn net.Conn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (t *svrTransHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
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

func (t *svrTransHandler) SetInvokeHandleFunc(inkHdlFunc endpoint.Endpoint) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) SetPipeline(p *remote.TransPipeline) { _ = "STUB: not implemented"; return }

func (t *svrTransHandler) writeErrorReplyIfNeeded(
	ctx context.Context, recvMsg remote.Message, conn net.Conn, err error, ri rpcinfo.RPCInfo, doOnMessage, connReset bool,
) (shouldCloseConn bool) {
	_ = "STUB: not implemented"
	return false
}

func (t *svrTransHandler) startTracer(ctx context.Context, ri rpcinfo.RPCInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (t *svrTransHandler) finishTracer(ctx context.Context, ri rpcinfo.RPCInfo, err error, conn net.Conn, panicErr interface{}) {
	_ = "STUB: not implemented"
	return
}

func (t *svrTransHandler) startProfiler(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (t *svrTransHandler) finishProfiler(ctx context.Context) { _ = "STUB: not implemented"; return }

func (t *svrTransHandler) GracefulShutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func getRemoteInfo(ri rpcinfo.RPCInfo, conn net.Conn) (string, net.Addr) {
	_ = "STUB: not implemented"
	return "", *new(net.Addr)
}
