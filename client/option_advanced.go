package client

import (
	"context"
	"crypto/tls"

	"github.com/cloudwego/kitex/internal/client"
	"github.com/cloudwego/kitex/pkg/acl"
	"github.com/cloudwego/kitex/pkg/diagnosis"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/proxy"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func WithHTTPConnection() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithClientBasicInfo(ebi *rpcinfo.EndpointBasicInfo) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDiagnosisService(ds diagnosis.Service) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithACLRules(rules ...acl.RejectFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithFirstMetaHandler(h remote.MetaHandler) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithMetaHandler(h remote.MetaHandler) client.Option {
	_ = "STUB: not implemented"
	return *new(client.Option)
}

func WithProxy(p proxy.ForwardProxy) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTransHandlerFactory(f remote.ClientTransHandlerFactory) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithDialer(d remote.Dialer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithConnPool(pool remote.ConnPool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithRetryContainer(rc *retry.Container) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGeneric(g generic.Generic) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCloseCallbacks(callback func() error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithErrorHandler(f func(context.Context, error) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithBoundHandler(h remote.BoundHandler) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithGRPCTLSConfig(tlsConfig *tls.Config) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
