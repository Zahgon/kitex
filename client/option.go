package client

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/internal/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/http"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/loadbalance/lbcache"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/pkg/warmup"
	"github.com/cloudwego/kitex/pkg/xds"
	"github.com/cloudwego/kitex/transport"
)

type Option = client.Option

type Options = client.Options

type UnaryOption = client.UnaryOption

type UnaryOptions = client.UnaryOptions

type StreamOption = client.StreamOption

type StreamOptions = client.StreamOptions

type TTHeaderStreamingOption = client.TTHeaderStreamingOption

type TTHeaderStreamingOptions = client.TTHeaderStreamingOptions

type Suite interface {
	Options() []Option
}

func WithTransportProtocol(tp transport.Protocol) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSuite(suite Suite) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMiddleware(mw endpoint.Middleware) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMiddlewareBuilder(mwb endpoint.MiddlewareBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithInstanceMW(mw endpoint.Middleware) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDestService(svr string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHostPorts(hostports ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithResolver(r discovery.Resolver) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHTTPResolver(r http.Resolver) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithShortConnection() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLongConnection(cfg connpool.IdleConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMuxConnection(connNum int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLogger(logger klog.FormatLogger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLoadBalancer(lb loadbalance.Loadbalancer, opts ...*lbcache.Options) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithRPCTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithConnectTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTimeoutProvider(p rpcinfo.TimeoutProvider) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTag(key, val string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTracer(c stats.Tracer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithStatsLevel(level stats.Level) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCodec(c remote.Codec) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPayloadCodec(c remote.PayloadCodec) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithConnReporterEnabled() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFailureRetry(p *retry.FailurePolicy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithBackupRequest(p *retry.BackupPolicy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMixedRetry(p *retry.MixedPolicy) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRetryMethodPolicies(mp map[string]retry.Policy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSpecifiedResultRetry(rr *retry.ShouldResultRetry) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithFallback(fb *fallback.Policy) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCircuitBreaker(s *circuitbreak.CBSuite) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithGRPCConnPoolSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCWriteBufferSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCReadBufferSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCInitialWindowSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCInitialConnWindowSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCMaxHeaderListSize(s uint32) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCKeepaliveParams(kp grpc.ClientKeepalive) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithGRPCReuseWriteBuffer(cfg grpc.ReuseWriteBufferConfig) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithWarmingUp(wuo *warmup.ClientOption) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithXDSSuite(suite xds.ClientSuite) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithContextBackup(backupHandler func(prev, cur context.Context) (ctx context.Context, backup bool)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func TailOption(opt Option) Option { _ = "STUB: not implemented"; return *new(Option) }
