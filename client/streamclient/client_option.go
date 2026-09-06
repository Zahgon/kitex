package streamclient

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/http"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/cloudwego/kitex/pkg/loadbalance/lbcache"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/pkg/warmup"
	"github.com/cloudwego/kitex/pkg/xds"
)

func WithSuite(suite client.Suite) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMiddleware(mw endpoint.Middleware) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMiddlewareBuilder(mwb endpoint.MiddlewareBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithInstanceMW(mw endpoint.Middleware) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDestService(svr string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHostPorts(hostPorts ...string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithResolver(r discovery.Resolver) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHTTPResolver(r http.Resolver) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLoadBalancer(lb loadbalance.Loadbalancer, opts ...*lbcache.Options) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithConnectTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTag(key, val string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTracer(c stats.Tracer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithStatsLevel(level stats.Level) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithPayloadCodec(c remote.PayloadCodec) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithConnReporterEnabled() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWarmingUp(wuo *warmup.ClientOption) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithXDSSuite(suite xds.ClientSuite) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithContextBackup(backupHandler func(prev, cur context.Context) (ctx context.Context, backup bool)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
