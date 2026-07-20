package logid

import (
	"context"
)

type keyTypeStreamLogID string

const ctxKeyStreamLogID keyTypeStreamLogID = "stream-log-id"

func NewCtxWithStreamLogID(ctx context.Context, logID string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetStreamLogID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

func GenerateStreamLogID(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
