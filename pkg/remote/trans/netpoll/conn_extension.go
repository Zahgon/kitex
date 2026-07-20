package netpoll

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func NewNetpollConnExtension() trans.Extension {
	_ = "STUB: not implemented"
	return *new(trans.Extension)
}

type netpollConnExtension struct{}

func (e *netpollConnExtension) SetReadTimeout(ctx context.Context, conn net.Conn, cfg rpcinfo.RPCConfig, role remote.RPCRole) {
	_ = "STUB: not implemented"
	return
}

func (e *netpollConnExtension) NewWriteByteBuffer(ctx context.Context, conn net.Conn, msg remote.Message) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (e *netpollConnExtension) NewReadByteBuffer(ctx context.Context, conn net.Conn, msg remote.Message) remote.ByteBuffer {
	_ = "STUB: not implemented"
	return *new(remote.ByteBuffer)
}

func (e *netpollConnExtension) ReleaseBuffer(buffer remote.ByteBuffer, err error) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *netpollConnExtension) IsTimeoutErr(err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *netpollConnExtension) IsRemoteClosedErr(err error) bool {
	_ = "STUB: not implemented"
	return false
}
