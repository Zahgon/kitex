package ttstream

import (
	"context"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

var _ remote.ClientStreamFactory = (*clientTransHandler)(nil)

func NewCliTransHandlerFactory(opts ...ClientHandlerOption) remote.ClientStreamFactory {
	_ = "STUB: not implemented"
	return *new(remote.ClientStreamFactory)
}

type clientTransHandler struct {
	transPool     transPool
	metaHandler   MetaFrameHandler
	headerHandler HeaderFrameWriteHandler
	traceCtl      *rpcinfo.TraceController
}

func (c clientTransHandler) NewStream(ctx context.Context, ri rpcinfo.RPCInfo) (streaming.ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(streaming.ClientStream), nil
}
