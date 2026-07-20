package thrift

import (
	"context"

	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/protocol/thrift/base"
)

const (
	structWrapLen = 4
)

type MessageReader interface {
	Read(ctx context.Context, method string, isClient bool, dataLen int, in bufiox.Reader) (interface{}, error)
}

type MessageWriter interface {
	Write(ctx context.Context, out bufiox.Writer, msg interface{}, method string, isClient bool, requestBase *base.Base) error
}
