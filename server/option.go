package server

import (
	"context"
	"net"
	"time"

	internal_server "github.com/cloudwego/kitex/internal/server"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/limiter"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/pkg/streaming"
)

type Option = internal_server.Option

type Options = internal_server.Options

type Suite interface {
	Options() []Option
}

type UnaryOption = internal_server.UnaryOption

type UnaryOptions = internal_server.UnaryOptions

type StreamOption = internal_server.StreamOption

type StreamOptions = internal_server.StreamOptions

type TTHeaderStreamingOption = remote.TTHeaderStreamingOption

func WithSuite(suite Suite) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMuxTransport() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMiddleware(mw endpoint.Middleware) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMiddlewareBuilder(mwb endpoint.MiddlewareBuilder, funcName ...string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithReadWriteTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLogger(logger klog.FormatLogger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithExitWaitTime(timeout time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxConnIdleTime(timeout time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLimit(lim *limit.Option) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithConnectionLimiter(conLimit limiter.ConcurrencyLimiter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithQPSLimiter(qpsLimit limiter.RateLimiter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTracer(c stats.Tracer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithStatsLevel(level stats.Level) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithServiceAddr(addr net.Addr) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCodec(c remote.Codec) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPayloadCodec(c remote.PayloadCodec) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRegistry(r registry.Registry) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRegistryInfo(info *registry.Info) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCWriteBufferSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCReadBufferSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCInitialWindowSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCInitialConnWindowSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCKeepaliveParams(kp grpc.ServerKeepalive) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithGRPCKeepaliveEnforcementPolicy(kep grpc.EnforcementPolicy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithGRPCMaxConcurrentStreams(n uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCMaxHeaderListSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCUnknownServiceHandler(f func(ctx context.Context, methodName string, stream streaming.Stream) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithGRPCReuseWriteBuffer(cfg grpc.ReuseWriteBufferConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithConcurrencyLimiter(conLimit limiter.ConcurrencyLimiter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithContextBackup(enable, async bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRefuseTrafficWithoutServiceName() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEnableContextTimeout(enable bool) Option { _ = "STUB: not implemented"; return *new(Option) }
