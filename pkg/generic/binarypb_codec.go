package generic

import (
	"github.com/cloudwego/kitex/pkg/generic/proto"
)

type binaryPbCodec struct {
	svcName      string
	packageName  string
	readerWriter *proto.RawReaderWriter
}

func newBinaryPbCodec(svcName, packageName string) *binaryPbCodec {
	_ = "STUB: not implemented"
	return nil
}

func (c *binaryPbCodec) getMessageReaderWriter() interface{} { _ = "STUB: not implemented"; return nil }

func (c *binaryPbCodec) Name() string { _ = "STUB: not implemented"; return "" }
