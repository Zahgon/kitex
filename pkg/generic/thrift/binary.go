package thrift

import (
	"context"

	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/protocol/thrift/base"
)

type WriteBinary struct{}

func NewWriteBinary() *WriteBinary { _ = "STUB: not implemented"; return nil }

func (w *WriteBinary) Write(ctx context.Context, out bufiox.Writer, msg interface{}, method string, isClient bool, requestBase *base.Base) error {
	_ = "STUB: not implemented"
	return nil
}
