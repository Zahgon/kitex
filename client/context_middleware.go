package client

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

type ctxMWChainKey struct{}

func WithContextMiddlewares(ctx context.Context, mws ...endpoint.Middleware) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func getContextMiddleware(ctx context.Context) endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

func contextMW(next endpoint.Endpoint) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}
