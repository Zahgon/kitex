package streamclient

import (
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func WithRecvMiddleware(mw endpoint.RecvMiddleware) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithRecvMiddlewareBuilder(mwb endpoint.RecvMiddlewareBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSendMiddleware(mw endpoint.SendMiddleware) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithSendMiddlewareBuilder(mwb endpoint.SendMiddlewareBuilder) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
