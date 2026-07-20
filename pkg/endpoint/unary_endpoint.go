package endpoint

import "context"

type UnaryEndpoint Endpoint

type UnaryMiddleware func(next UnaryEndpoint) UnaryEndpoint

type UnaryMiddlewareBuilder func(ctx context.Context) UnaryMiddleware

func UnaryChain(mws ...UnaryMiddleware) UnaryMiddleware {
	_ = "STUB: not implemented"
	return *new(UnaryMiddleware)
}

func (mw UnaryMiddleware) ToMiddleware() Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}
