package remote

import (
	"context"
)

var (
	_ MetaHandler          = (*customMetaHandler)(nil)
	_ StreamingMetaHandler = (*customMetaHandler)(nil)
)

type CustomMetaHandlerOption func(*customMetaHandler)

func NewCustomMetaHandler(opts ...CustomMetaHandlerOption) MetaHandler {
	_ = "STUB: not implemented"
	return *new(MetaHandler)
}

func WithOnReadStream(fn func(ctx context.Context) (context.Context, error)) CustomMetaHandlerOption {
	_ = "STUB: not implemented"
	return *new(CustomMetaHandlerOption)
}

func WithOnConnectStream(fn func(ctx context.Context) (context.Context, error)) CustomMetaHandlerOption {
	_ = "STUB: not implemented"
	return *new(CustomMetaHandlerOption)
}

func WithWriteMeta(fn func(ctx context.Context, msg Message) (context.Context, error)) CustomMetaHandlerOption {
	_ = "STUB: not implemented"
	return *new(CustomMetaHandlerOption)
}

func WithReadMeta(fn func(ctx context.Context, msg Message) (context.Context, error)) CustomMetaHandlerOption {
	_ = "STUB: not implemented"
	return *new(CustomMetaHandlerOption)
}

type customMetaHandler struct {
	writeMeta       func(ctx context.Context, msg Message) (context.Context, error)
	readMeta        func(ctx context.Context, msg Message) (context.Context, error)
	onReadStream    func(ctx context.Context) (context.Context, error)
	onConnectStream func(ctx context.Context) (context.Context, error)
}

func (c *customMetaHandler) WriteMeta(ctx context.Context, msg Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (c *customMetaHandler) ReadMeta(ctx context.Context, msg Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (c *customMetaHandler) OnConnectStream(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (c *customMetaHandler) OnReadStream(ctx context.Context) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
