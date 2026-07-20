package rpcinfo

import (
	"context"
)

type ctxRPCInfoKeyType struct{}

var ctxRPCInfoKey ctxRPCInfoKeyType

func NewCtxWithRPCInfo(ctx context.Context, ri RPCInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetRPCInfo(ctx context.Context) RPCInfo { _ = "STUB: not implemented"; return *new(RPCInfo) }

func PutRPCInfo(ri RPCInfo) { _ = "STUB: not implemented"; return }

func FreezeRPCInfo(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
