package generic

import (
	"sync/atomic"

	"github.com/cloudwego/dynamicgo/conv"

	"github.com/cloudwego/kitex/pkg/generic/descriptor"
	"github.com/cloudwego/kitex/pkg/generic/thrift"
)

var _ Closer = &jsonThriftCodec{}

type jsonThriftCodec struct {
	svcDsc                 atomic.Value
	provider               DescriptorProvider
	binaryWithBase64       bool
	dynamicgoEnabled       bool
	convOpts               conv.Options
	convOptsWithThriftBase conv.Options
	convOptsWithException  conv.Options
	svcName                atomic.Value
	combineService         atomic.Value
	readerWriter           atomic.Value
}

func newJsonThriftCodec(p DescriptorProvider, opts *Options) *jsonThriftCodec {
	_ = "STUB: not implemented"
	return nil
}

func (c *jsonThriftCodec) update() { _ = "STUB: not implemented"; return }

func (c *jsonThriftCodec) updateMessageReaderWriter() (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *jsonThriftCodec) configureMessageReaderWriter(svc *descriptor.ServiceDescriptor) {
	_ = "STUB: not implemented"
	return
}

func (c *jsonThriftCodec) getMessageReaderWriter() interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *jsonThriftCodec) configureJSONWriter(writer *thrift.WriteJSON) {
	_ = "STUB: not implemented"
	return
}

func (c *jsonThriftCodec) configureJSONReader(reader *thrift.ReadJSON) {
	_ = "STUB: not implemented"
	return
}

func (c *jsonThriftCodec) getMethod(method string) (Method, error) {
	_ = "STUB: not implemented"
	return *new(Method), nil
}

func (c *jsonThriftCodec) Name() string { _ = "STUB: not implemented"; return "" }

func (c *jsonThriftCodec) Close() error { _ = "STUB: not implemented"; return nil }
