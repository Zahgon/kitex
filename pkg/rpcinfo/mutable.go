package rpcinfo

import (
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/transport"
)

type MutableEndpointInfo interface {
	SetServiceName(service string) error
	SetMethod(method string) error
	SetAddress(addr net.Addr) error
	SetTag(key, value string) error
	ImmutableView() EndpointInfo
	Reset()
	ResetFromBasicInfo(bi *EndpointBasicInfo)
}

type MutableRPCConfig interface {
	SetRPCTimeout(to time.Duration) error
	IsRPCTimeoutLocked() bool
	SetConnectTimeout(to time.Duration) error
	IsConnectTimeoutLocked() bool
	SetReadWriteTimeout(to time.Duration) error
	IsReadWriteTimeoutLocked() bool
	SetIOBufferSize(sz int) error
	SetTransportProtocol(tp transport.Protocol) error
	SetInteractionMode(mode InteractionMode) error
	LockConfig(bits int)
	Clone() MutableRPCConfig
	CopyFrom(from RPCConfig)
	ImmutableView() RPCConfig
	SetPayloadCodec(codec serviceinfo.PayloadCodec)

	SetStreamRecvTimeout(timeout time.Duration)
	SetStreamRecvTimeoutConfig(cfg streaming.TimeoutConfig)
}

type MutableRPCStats interface {
	SetSendSize(size uint64)
	SetRecvSize(size uint64)
	SetError(err error)
	SetPanicked(x interface{})
	SetLevel(level stats.Level)
	Reset()
	ImmutableView() RPCStats
	IncrSendSize(size uint64)
	IncrRecvSize(size uint64)
}
