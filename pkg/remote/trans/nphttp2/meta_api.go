package nphttp2

import (
	"context"

	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

func SetHeader(ctx context.Context, md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func SendHeader(ctx context.Context, md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func SetTrailer(ctx context.Context, md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func serverTransportStreamFromContext(ctx context.Context) streaming.Stream {
	_ = "STUB: not implemented"
	return *new(streaming.Stream)
}

type (
	headerKey  struct{}
	trailerKey struct{}
)

func GRPCHeader(ctx context.Context, md *metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GRPCTrailer(ctx context.Context, md *metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetHeaderMetadataFromCtx(ctx context.Context) *metadata.MD {
	_ = "STUB: not implemented"
	return nil
}

func GetTrailerMetadataFromCtx(ctx context.Context) *metadata.MD {
	_ = "STUB: not implemented"
	return nil
}

var unaryMetaEventHandler = rpcinfo.ClientStreamEventHandler{
	HandleStreamRecvHeaderEvent: unaryMetaHandleRecvHeaderEvent,
	HandleStreamFinishEvent:     unaryMetaHandleFinishEvent,
}

func unaryMetaHandleRecvHeaderEvent(ctx context.Context, ri rpcinfo.RPCInfo, evt rpcinfo.StreamRecvHeaderEvent) {
	_ = "STUB: not implemented"
	return
}

func unaryMetaHandleFinishEvent(ctx context.Context, ri rpcinfo.RPCInfo, evt rpcinfo.StreamFinishEvent) {
	_ = "STUB: not implemented"
	return
}

func isUnary(ri rpcinfo.RPCInfo) bool { _ = "STUB: not implemented"; return false }
