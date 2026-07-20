package server

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/acl"
	"github.com/cloudwego/kitex/pkg/diagnosis"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/limiter"
	"github.com/cloudwego/kitex/pkg/profiler"
	"github.com/cloudwego/kitex/pkg/proxy"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func WithServerBasicInfo(ebi *rpcinfo.EndpointBasicInfo) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDiagnosisService(ds diagnosis.Service) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithACLRules(rules ...acl.RejectFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMetaHandler(h remote.MetaHandler) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProxy(p proxy.ReverseProxy) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTransHandlerFactory(f remote.ServerTransHandlerFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTransServerFactory(f remote.TransServerFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithLimitReporter(r limiter.LimitReporter) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithGeneric(g generic.Generic) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithErrorHandler(f func(context.Context, error) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithBoundHandler(h remote.BoundHandler) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithExitSignal(f func() <-chan error) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithListener(ln net.Listener) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReusePort(reuse bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProfiler(pc profiler.Profiler) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProfilerTransInfoTagging(tagging remote.TransInfoTagging) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithProfilerMessageTagging(tagging remote.MessageTagging) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
