package generic

import (
	"sync/atomic"

	"github.com/cloudwego/kitex/pkg/generic/descriptor"
	"github.com/cloudwego/kitex/pkg/generic/thrift"
)

var _ Closer = &mapThriftCodec{}

type mapThriftCodec struct {
	svcDsc                  atomic.Value
	provider                DescriptorProvider
	forJSON                 bool
	binaryWithBase64        bool
	binaryWithByteSlice     bool
	setFieldsForEmptyStruct uint8
	svcName                 atomic.Value
	combineService          atomic.Value
	readerWriter            atomic.Value
}

func newMapThriftCodec(p DescriptorProvider, forJSON bool) *mapThriftCodec {
	_ = "STUB: not implemented"
	return nil
}

func (c *mapThriftCodec) update() { _ = "STUB: not implemented"; return }

func (c *mapThriftCodec) updateMessageReaderWriter() (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *mapThriftCodec) configureMessageReaderWriter(svc *descriptor.ServiceDescriptor) {
	_ = "STUB: not implemented"
	return
}

func (c *mapThriftCodec) getMessageReaderWriter() interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *mapThriftCodec) configureStructWriter(writer *thrift.WriteStruct) {
	_ = "STUB: not implemented"
	return
}

func (c *mapThriftCodec) configureStructReader(reader *thrift.ReadStruct) {
	_ = "STUB: not implemented"
	return
}

func (c *mapThriftCodec) getMethod(method string) (Method, error) {
	_ = "STUB: not implemented"
	return *new(Method), nil
}

func (c *mapThriftCodec) Name() string { _ = "STUB: not implemented"; return "" }

func (c *mapThriftCodec) Close() error { _ = "STUB: not implemented"; return nil }
