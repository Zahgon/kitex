package server

import (
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func WithUnaryOptions(opts ...UnaryOption) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithUnaryMiddleware(mw endpoint.UnaryMiddleware) UnaryOption {
	_ = "STUB: not implemented"
	return *new(UnaryOption)
}

func WithUnaryMiddlewareBuilder(mwb endpoint.UnaryMiddlewareBuilder) UnaryOption {
	_ = "STUB: not implemented"
	return *new(UnaryOption)
}
