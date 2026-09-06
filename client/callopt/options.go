package callopt

import (
	"strings"
	"sync"
	"time"

	"github.com/cloudwego/kitex/internal/client"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/http"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/rpcinfo/remoteinfo"
)

var callOptionsPool = sync.Pool{
	New: newOptions,
}

type CallOptions struct {
	configs      rpcinfo.MutableRPCConfig
	svr          remoteinfo.RemoteInfo
	locks        *client.ConfigLocks
	httpResolver http.Resolver

	RetryPolicy             retry.Policy
	Fallback                *fallback.Policy
	CompressorName          string
	StreamOptions           client.StreamOptions
	BinaryGenericIDLService string
}

func newOptions() interface{} { _ = "STUB: not implemented"; return nil }

func (co *CallOptions) Recycle() { _ = "STUB: not implemented"; return }

type Option struct {
	f func(o *CallOptions, di *strings.Builder)
}

func (o Option) F() func(o *CallOptions, di *strings.Builder) {
	_ = "STUB: not implemented"
	return nil
}

func NewOption(f func(o *CallOptions, di *strings.Builder)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithHostPort(hostport string) Option { _ = "STUB: not implemented"; return *new(Option) }

func setInstance(svr remoteinfo.RemoteInfo, hostport string) error {
	_ = "STUB: not implemented"
	return nil
}

func WithURL(url string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithHTTPHost(host string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRPCTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithConnectTimeout(d time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTag(key, val string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRetryPolicy(p retry.Policy) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFallback(fb *fallback.Policy) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCCompressor(compressorName string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithBinaryGenericIDLService(svcName string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func Apply(cos []Option, cfg rpcinfo.MutableRPCConfig, svr remoteinfo.RemoteInfo, locks *client.ConfigLocks, httpResolver http.Resolver) (string, *CallOptions) {
	_ = "STUB: not implemented"
	return "", nil
}
