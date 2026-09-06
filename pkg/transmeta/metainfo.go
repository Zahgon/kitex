package transmeta

import (
	"context"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
)

var (
	MetainfoClientHandler = new(metainfoClientHandler)
	MetainfoServerHandler = new(metainfoServerHandler)

	_ remote.MetaHandler          = MetainfoClientHandler
	_ remote.StreamingMetaHandler = MetainfoClientHandler
	_ remote.MetaHandler          = MetainfoServerHandler
	_ remote.StreamingMetaHandler = MetainfoServerHandler
)

type metainfoClientHandler struct{}

func (mi *metainfoClientHandler) OnConnectStream(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (mi *metainfoClientHandler) OnReadStream(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (mi *metainfoClientHandler) WriteMeta(ctx context.Context, sendMsg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (mi *metainfoClientHandler) ReadMeta(ctx context.Context, recvMsg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

type metainfoServerHandler struct{}

func (mi *metainfoServerHandler) ReadMeta(ctx context.Context, recvMsg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (mi *metainfoServerHandler) OnConnectStream(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (mi *metainfoServerHandler) OnReadStream(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func addStreamIDToContext(ctx context.Context, md metadata.MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (mi *metainfoServerHandler) WriteMeta(ctx context.Context, sendMsg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
