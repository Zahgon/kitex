package nphttp2

import (
	"container/list"
	"context"
	"net"
	"sync"
	"time"

	igeneric "github.com/cloudwego/kitex/internal/generic"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/codec"
	grpcTransport "github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

var streamingBidirectionalCtx = igeneric.WithGenericStreamingMode(context.Background(), serviceinfo.StreamingBidirectional)

type svrTransHandlerFactory struct{}

func NewSvrTransHandlerFactory() remote.ServerTransHandlerFactory {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandlerFactory)
}

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

	mu sync.Mutex

	li *list.List
}

var prefaceReadAtMost = func() int {

	if 2*codec.Size32 < grpcTransport.ClientPrefaceLen {
		return 2 * codec.Size32
	}
	return grpcTransport.ClientPrefaceLen
}()

func (t *svrTransHandler) ProtocolMatch(ctx context.Context, conn net.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *svrTransHandler) Write(ctx context.Context, conn net.Conn, msg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) Read(ctx context.Context, conn net.Conn, msg remote.Message) (nctx context.Context, err error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (t *svrTransHandler) OnRead(ctx context.Context, conn net.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *svrTransHandler) handleFunc(s *grpcTransport.Stream, svrTrans *SvrTrans, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

//nolint:staticcheck // SA1029: consts.CtxKeyMethod has been used and we just follow it

func invokeStreamUnaryHandler(ctx context.Context, st streaming.ServerStream, mi serviceinfo.MethodInfo,
	handler endpoint.Endpoint, ri rpcinfo.RPCInfo,
) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func getPayloadCodecFromContentType(ct string) serviceinfo.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(serviceinfo.PayloadCodec)
}

func (t *svrTransHandler) OnMessage(ctx context.Context, args, result remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

type svrTransKey int

const (
	ctxKeySvrTransport svrTransKey = 1

	defaultGraceTime time.Duration = 5 * time.Second

	defaultMaxPollTime = 50 * time.Millisecond
)

type SvrTrans struct {
	tr   grpcTransport.ServerTransport
	pool *sync.Pool
	elem *list.Element

	handlerNum int32
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

func (t *svrTransHandler) GracefulShutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func parseGraceAndPollTime(ctx context.Context) (graceTime, pollTime time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration)
}

func (t *svrTransHandler) startTracer(ctx context.Context, ri rpcinfo.RPCInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (t *svrTransHandler) finishTracer(ctx context.Context, ri rpcinfo.RPCInfo, err error, panicErr interface{}) {
	_ = "STUB: not implemented"
	return
}
