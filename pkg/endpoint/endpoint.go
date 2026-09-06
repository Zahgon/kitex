package endpoint

import "context"

type Endpoint func(ctx context.Context, req, resp interface{}) (err error)

type Middleware func(next Endpoint) Endpoint

type MiddlewareBuilder func(ctx context.Context) Middleware

func Chain(mws ...Middleware) Middleware { _ = "STUB: not implemented"; return *new(Middleware) }

func Build(mws []Middleware) Middleware { _ = "STUB: not implemented"; return *new(Middleware) }

func (mw Middleware) ToUnaryMiddleware() UnaryMiddleware {
	_ = "STUB: not implemented"
	return *new(UnaryMiddleware)
}

func DummyMiddleware(next Endpoint) Endpoint { _ = "STUB: not implemented"; return *new(Endpoint) }

func DummyEndpoint(ctx context.Context, req, resp interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type mwCtxKeyType int

const (
	CtxEventBusKey mwCtxKeyType = iota
	CtxEventQueueKey
)
