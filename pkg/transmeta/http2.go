package transmeta

import (
	"context"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var ClientHTTP2Handler = &clientHTTP2Handler{}

type clientHTTP2Handler struct{}

var (
	_ remote.MetaHandler          = ClientHTTP2Handler
	_ remote.StreamingMetaHandler = ClientHTTP2Handler
)

func metaAppend(m metadata.MD, k, v string) { _ = "STUB: not implemented"; return }

func (*clientHTTP2Handler) OnConnectStream(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (*clientHTTP2Handler) OnReadStream(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (ch *clientHTTP2Handler) WriteMeta(ctx context.Context, msg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (ch *clientHTTP2Handler) ReadMeta(ctx context.Context, msg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

var ServerHTTP2Handler = &serverHTTP2Handler{}

type serverHTTP2Handler struct{}

func (*serverHTTP2Handler) OnConnectStream(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (*serverHTTP2Handler) OnReadStream(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (sh *serverHTTP2Handler) WriteMeta(ctx context.Context, msg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (sh *serverHTTP2Handler) ReadMeta(ctx context.Context, msg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func isGRPC(ri rpcinfo.RPCInfo) bool { _ = "STUB: not implemented"; return false }
