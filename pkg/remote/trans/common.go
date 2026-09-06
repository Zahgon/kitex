package trans

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var (
	readMoreTimeout = 5 * time.Millisecond

	ErrRemoteClosed = errors.New("remote connection closed")
)

type RemoteClosedSource int

const (
	RemoteClosedByExtension RemoteClosedSource = iota

	RemoteClosedByConnectionState
)

func (s RemoteClosedSource) String() string { _ = "STUB: not implemented"; return "" }

type RemoteClosedError struct {
	Source RemoteClosedSource
	Cause  error
}

func (e *RemoteClosedError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *RemoteClosedError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e *RemoteClosedError) Is(target error) bool { _ = "STUB: not implemented"; return false }

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

func IsRemoteClosedErr(ext Extension, err error, conn net.Conn) *RemoteClosedError {
	_ = "STUB: not implemented"
	return nil
}

type MuxEnabledFlag interface {
	MuxEnabled() bool
}
