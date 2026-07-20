package transmeta

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

const (
	framedTransportType   = "framed"
	unframedTransportType = "unframed"

	bizStatus  = "biz-status"
	bizMessage = "biz-message"
	bizExtra   = "biz-extra"
)

var (
	ClientTTHeaderHandler remote.MetaHandler = &clientTTHeaderHandler{}
	ServerTTHeaderHandler remote.MetaHandler = &serverTTHeaderHandler{}
)

type clientTTHeaderHandler struct{}

func (ch *clientTTHeaderHandler) WriteMeta(ctx context.Context, msg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func getIDLSvcName(ri rpcinfo.RPCInfo) string { _ = "STUB: not implemented"; return "" }

func (ch *clientTTHeaderHandler) ReadMeta(ctx context.Context, msg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func ParseBizStatusErr(strInfo map[string]string) (kerrors.BizStatusErrorIface, error) {
	_ = "STUB: not implemented"
	return *new(kerrors.BizStatusErrorIface), nil
}

type serverTTHeaderHandler struct{}

func (sh *serverTTHeaderHandler) ReadMeta(ctx context.Context, msg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (sh *serverTTHeaderHandler) WriteMeta(ctx context.Context, msg remote.Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func isTTHeader(msg remote.Message) bool { _ = "STUB: not implemented"; return false }
