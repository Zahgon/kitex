package rpcinfo

import (
	"sync"
	"time"

	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/transport"
)

var (
	_             MutableRPCConfig = &rpcConfig{}
	_             RPCConfig        = &rpcConfig{}
	rpcConfigPool sync.Pool
)

var (
	defaultRPCTimeout       = time.Duration(0)
	defaultConnectTimeout   = time.Millisecond * 50
	defaultReadWriteTimeout = time.Second * 5
	defaultBufferSize       = 4096
	defaultInteractionMode  = PingPong
)

const (
	BitRPCTimeout = 1 << iota
	BitConnectTimeout
	BitReadWriteTimeout
	BitIOBufferSize
)

type InteractionMode int32

const (
	PingPong  InteractionMode = 0
	Oneway    InteractionMode = 1
	Streaming InteractionMode = 2
)

type rpcConfig struct {
	readOnlyMask      int
	rpcTimeout        time.Duration
	connectTimeout    time.Duration
	readWriteTimeout  time.Duration
	ioBufferSize      int
	transportProtocol transport.Protocol
	interactionMode   InteractionMode
	payloadCodec      serviceinfo.PayloadCodec

	streamRecvTimeout       time.Duration
	streamRecvTimeoutConfig streaming.TimeoutConfig
}

func init() {
	rpcConfigPool.New = newRPCConfig
}

func newRPCConfig() interface{} { _ = "STUB: not implemented"; return nil }

func (r *rpcConfig) LockConfig(bits int) { _ = "STUB: not implemented"; return }

func (r *rpcConfig) SetRPCTimeout(to time.Duration) error { _ = "STUB: not implemented"; return nil }

func (r *rpcConfig) IsRPCTimeoutLocked() bool { _ = "STUB: not implemented"; return false }

func (r *rpcConfig) SetConnectTimeout(to time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rpcConfig) IsConnectTimeoutLocked() bool { _ = "STUB: not implemented"; return false }

func (r *rpcConfig) SetReadWriteTimeout(to time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rpcConfig) IsReadWriteTimeoutLocked() bool { _ = "STUB: not implemented"; return false }

func (r *rpcConfig) SetIOBufferSize(sz int) error { _ = "STUB: not implemented"; return nil }

func (r *rpcConfig) ImmutableView() RPCConfig { _ = "STUB: not implemented"; return *new(RPCConfig) }

func (r *rpcConfig) RPCTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *rpcConfig) ConnectTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *rpcConfig) ReadWriteTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *rpcConfig) IOBufferSize() int { _ = "STUB: not implemented"; return 0 }

func (r *rpcConfig) TransportProtocol() transport.Protocol {
	_ = "STUB: not implemented"
	return *new(transport.Protocol)
}

func (r *rpcConfig) SetTransportProtocol(tp transport.Protocol) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rpcConfig) SetInteractionMode(mode InteractionMode) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rpcConfig) InteractionMode() InteractionMode {
	_ = "STUB: not implemented"
	return *new(InteractionMode)
}

func (r *rpcConfig) SetPayloadCodec(codec serviceinfo.PayloadCodec) {
	_ = "STUB: not implemented"
	return
}

func (r *rpcConfig) PayloadCodec() serviceinfo.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(serviceinfo.PayloadCodec)
}

func (r *rpcConfig) SetStreamRecvTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

func (r *rpcConfig) StreamRecvTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (r *rpcConfig) SetStreamRecvTimeoutConfig(cfg streaming.TimeoutConfig) {
	_ = "STUB: not implemented"
	return
}

func (r *rpcConfig) StreamRecvTimeoutConfig() streaming.TimeoutConfig {
	_ = "STUB: not implemented"
	return *new(streaming.TimeoutConfig)
}

func (r *rpcConfig) Clone() MutableRPCConfig {
	_ = "STUB: not implemented"
	return *new(MutableRPCConfig)
}

func (r *rpcConfig) CopyFrom(from RPCConfig) { _ = "STUB: not implemented"; return }

func (r *rpcConfig) initialize() { _ = "STUB: not implemented"; return }

func (r *rpcConfig) Recycle() { _ = "STUB: not implemented"; return }

func NewRPCConfig() RPCConfig { _ = "STUB: not implemented"; return *new(RPCConfig) }
