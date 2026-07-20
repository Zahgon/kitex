package gonet

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func NewGonetExtension() trans.Extension { _ = "STUB: not implemented"; return *new(trans.Extension) }

type gonetConnExtension struct{}

func (e *gonetConnExtension) SetReadTimeout(ctx context.Context, conn net.Conn, cfg rpcinfo.RPCConfig, role remote.RPCRole) {
	_ = "STUB: not implemented"
	return
}

func (e *gonetConnExtension) NewWriteByteBuffer(ctx context.Context, conn net.Conn, msg remote.Message) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (e *gonetConnExtension) NewReadByteBuffer(ctx context.Context, conn net.Conn, msg remote.Message) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (e *gonetConnExtension) ReleaseBuffer(buffer remote.ByteBuffer, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *gonetConnExtension) IsTimeoutErr(err error) bool { _ = "STUB: not implemented"; return false }

func (e *gonetConnExtension) IsRemoteClosedErr(err error) bool {
	_ = "STUB: not implemented"
	return false
}
