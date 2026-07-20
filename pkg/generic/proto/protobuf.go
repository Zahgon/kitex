package proto

import (
	"context"
)

type MessageReader interface {
	Read(ctx context.Context, method string, isClient bool, actualMsgBuf []byte) (interface{}, error)
}

type MessageWriter interface {
	Write(ctx context.Context, msg interface{}, method string, isClient bool) (interface{}, error)
}
