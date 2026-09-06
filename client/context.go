package client

import (
	"context"

	"github.com/cloudwego/kitex/client/callopt"
)

type ctxKeyType int

const (
	ctxCallOptionKey ctxKeyType = iota
	ctxCallOptionInfoKey
)

func NewCtxWithCallOptions(ctx context.Context, opts []callopt.Option) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func CallOptionsFromCtx(ctx context.Context) (res []callopt.Option) {
	_ = "STUB: not implemented"
	return nil
}

func CallOptionInfoFromCtx(ctx context.Context) (res string) { _ = "STUB: not implemented"; return "" }
