package streamclient

import (
	"context"

	"github.com/cloudwego/kitex/pkg/acl"
	"github.com/cloudwego/kitex/pkg/diagnosis"
	"github.com/cloudwego/kitex/pkg/proxy"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

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

func WithMetaHandler(h remote.MetaHandler) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithProxy(p proxy.ForwardProxy) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithDialer(d remote.Dialer) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCloseCallbacks(callback func() error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithErrorHandler(f func(context.Context, error) error) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
