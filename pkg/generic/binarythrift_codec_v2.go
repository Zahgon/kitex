package generic

import (
	"github.com/cloudwego/kitex/pkg/generic/thrift"
)

type binaryThriftCodecV2 struct {
	svcName      string
	readerWriter *thrift.RawReaderWriter
}

func newBinaryThriftCodecV2(svcName string) *binaryThriftCodecV2 {
	_ = "STUB: not implemented"
	return nil
}

func (c *binaryThriftCodecV2) Name() string { _ = "STUB: not implemented"; return "" }

func (c *binaryThriftCodecV2) getMessageReaderWriter() interface{} {
	_ = "STUB: not implemented"
	return nil
}
