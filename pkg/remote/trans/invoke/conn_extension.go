package invoke

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func newIvkConnExtension() trans.Extension { _ = "STUB: not implemented"; return *new(trans.Extension) }

type ivkConnExtension struct{}

func (e *ivkConnExtension) SetReadTimeout(ctx context.Context, conn net.Conn, cfg rpcinfo.RPCConfig, role remote.RPCRole) {
	_ = "STUB: not implemented"
	return
}

func (e *ivkConnExtension) NewWriteByteBuffer(ctx context.Context, conn net.Conn, msg remote.Message) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (e *ivkConnExtension) NewReadByteBuffer(ctx context.Context, conn net.Conn, msg remote.Message) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (e *ivkConnExtension) ReleaseBuffer(remote.ByteBuffer, error) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ivkConnExtension) IsTimeoutErr(err error) bool { _ = "STUB: not implemented"; return false }

func (e *ivkConnExtension) IsRemoteClosedErr(err error) bool {
	_ = "STUB: not implemented"
	return false
}
