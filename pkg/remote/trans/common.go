package trans

import (
	"context"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var readMoreTimeout = 5 * time.Millisecond

type Extension interface {
	SetReadTimeout(ctx context.Context, conn net.Conn, cfg rpcinfo.RPCConfig, role remote.RPCRole)
	NewWriteByteBuffer(ctx context.Context, conn net.Conn, msg remote.Message) remote.ByteBuffer
	NewReadByteBuffer(ctx context.Context, conn net.Conn, msg remote.Message) remote.ByteBuffer
	ReleaseBuffer(remote.ByteBuffer, error) error
	IsTimeoutErr(error) bool

	IsRemoteClosedErr(error) bool
}

func GetReadTimeout(cfg rpcinfo.RPCConfig) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type MuxEnabledFlag interface {
	MuxEnabled() bool
}
