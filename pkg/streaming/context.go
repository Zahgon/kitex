package streaming

import (
	"context"
)

type (
	streamKey struct{}
)

func NewCtxWithStream(ctx context.Context, stream Stream) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetStream(ctx context.Context) Stream { _ = "STUB: not implemented"; return *new(Stream) }
