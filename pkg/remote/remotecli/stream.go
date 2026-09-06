package remotecli

import (
	"context"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

func NewStream(ctx context.Context, ri rpcinfo.RPCInfo, handler remote.ClientTransHandler, opt *remote.ClientOption) (streaming.ClientStream, *StreamConnManager, error) {
	_ = "STUB: not implemented"
	return *new(streaming.ClientStream), nil, nil
}

func NewStreamConnManager(cr ConnReleaser) *StreamConnManager {
	_ = "STUB: not implemented"
	return nil
}

type StreamConnManager struct {
	ConnReleaser
}

func (scm *StreamConnManager) ReleaseConn(err error, ri rpcinfo.RPCInfo) {
	_ = "STUB: not implemented"
	return
}
