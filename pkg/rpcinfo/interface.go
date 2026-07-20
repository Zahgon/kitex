package rpcinfo

import (
	"context"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/transport"
)

type EndpointInfo interface {
	ServiceName() string
	Method() string
	Address() net.Addr
	Tag(key string) (value string, exist bool)
	DefaultTag(key, def string) string
}

type RPCStats interface {
	Record(ctx context.Context, event stats.Event, status stats.Status, info string)
	SendSize() uint64

	LastSendSize() uint64
	RecvSize() uint64

	LastRecvSize() uint64
	Error() error
	Panicked() (bool, interface{})
	GetEvent(event stats.Event) Event
	Level() stats.Level
	CopyForRetry() RPCStats
}

type Event interface {
	Event() stats.Event
	Status() stats.Status
	Info() string
	Time() time.Time
	IsNil() bool
}

type Timeouts interface {
	RPCTimeout() time.Duration
	ConnectTimeout() time.Duration
	ReadWriteTimeout() time.Duration
}

type TimeoutProvider interface {
	Timeouts(ri RPCInfo) Timeouts
}

type StreamConfig interface {
	StreamRecvTimeout() time.Duration
	StreamRecvTimeoutConfig() streaming.TimeoutConfig
}

type RPCConfig interface {
	Timeouts
	StreamConfig
	IOBufferSize() int
	TransportProtocol() transport.Protocol
	InteractionMode() InteractionMode
	PayloadCodec() serviceinfo.PayloadCodec
}

type Invocation interface {
	PackageName() string
	ServiceName() string
	MethodName() string
	MethodInfo() serviceinfo.MethodInfo
	StreamingMode() serviceinfo.StreamingMode
	SeqID() int32
	BizStatusErr() kerrors.BizStatusErrorIface
	Extra(key string) interface{}
}

type RPCInfo interface {
	From() EndpointInfo
	To() EndpointInfo
	Invocation() Invocation
	Config() RPCConfig
	Stats() RPCStats
}
