package remote

import (
	"context"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/profiler"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/pkg/utils"
)

type Option struct {
	Outbounds []OutboundHandler

	Inbounds []InboundHandler

	StreamingMetaHandlers []StreamingMetaHandler
}

func (o *Option) PrependBoundHandler(h BoundHandler) { _ = "STUB: not implemented"; return }

func (o *Option) AppendBoundHandler(h BoundHandler) { _ = "STUB: not implemented"; return }

type ServerOption struct {
	SvcSearcher ServiceSearcher

	TransServerFactory TransServerFactory

	SvrHandlerFactory ServerTransHandlerFactory

	Codec Codec

	PayloadCodec PayloadCodec

	Listener net.Listener

	Address net.Addr

	ReusePort bool

	ExitWaitTime time.Duration

	AcceptFailedDelayTime time.Duration

	MaxConnectionIdleTime time.Duration

	ReadWriteTimeout time.Duration

	InitOrResetRPCInfoFunc func(rpcinfo.RPCInfo, net.Addr) rpcinfo.RPCInfo

	TracerCtl *rpcinfo.TraceController

	Profiler                 profiler.Profiler
	ProfilerTransInfoTagging TransInfoTagging
	ProfilerMessageTagging   MessageTagging

	GRPCCfg *grpc.ServerConfig

	GRPCUnknownServiceHandler func(ctx context.Context, method string, stream streaming.Stream) error

	TTHeaderStreamingOptions TTHeaderStreamingOptions

	Option

	CompatibleMiddlewareForUnary bool
}

type ClientOption struct {
	SvcInfo *serviceinfo.ServiceInfo

	CliHandlerFactory ClientTransHandlerFactory

	Codec Codec

	PayloadCodec PayloadCodec

	ConnPool ConnPool

	Dialer Dialer

	Option

	EnableConnPoolReporter bool

	GRPCStreamingCliHandlerFactory ClientTransHandlerFactory
	GRPCStreamingConnPool          ConnPool

	TTHeaderStreamingCliHandlerFactory ClientStreamFactory
}

type TTHeaderStreamingOption struct {
	F func(o *TTHeaderStreamingOptions, di *utils.Slice)
}

type TTHeaderStreamingOptions struct {
	TransportOptions []interface{}
}
