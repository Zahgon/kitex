package thrift

import (
	"context"

	"github.com/cloudwego/gopkg/bufiox"

	"github.com/cloudwego/kitex/pkg/remote"
)

const marshalThriftBufferSize = 1024

var defaultCodec = NewThriftCodec().(*thriftCodec)

func MarshalThriftData(ctx context.Context, codec remote.PayloadCodec, data interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c thriftCodec) marshalThriftData(ctx context.Context, data interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalThriftException(in bufiox.Reader) error { _ = "STUB: not implemented"; return nil }

func UnmarshalThriftData(ctx context.Context, codec remote.PayloadCodec, method string, buf []byte, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c thriftCodec) unmarshalThriftData(ctx context.Context, trans bufiox.Reader, data interface{}, dataLen int) error {
	_ = "STUB: not implemented"
	return nil
}
