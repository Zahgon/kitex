package generic

import (
	"sync/atomic"

	"github.com/cloudwego/dynamicgo/conv"
	dproto "github.com/cloudwego/dynamicgo/proto"

	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

var _ Closer = &jsonPbCodec{}

type jsonPbCodec struct {
	svcDsc           atomic.Value
	provider         PbDescriptorProviderDynamicGo
	opts             *Options
	convOpts         conv.Options
	dynamicgoEnabled bool
	svcName          atomic.Value
	packageName      atomic.Value
	combineService   atomic.Value
	readerWriter     atomic.Value
}

func newJsonPbCodec(p PbDescriptorProviderDynamicGo, opts *Options) *jsonPbCodec {
	_ = "STUB: not implemented"
	return nil
}

func (c *jsonPbCodec) update() { _ = "STUB: not implemented"; return }

func (c *jsonPbCodec) getMessageReaderWriter() interface{} { _ = "STUB: not implemented"; return nil }

func (c *jsonPbCodec) getMethod(method string) (Method, error) {
	_ = "STUB: not implemented"
	return *new(Method), nil
}

func (c *jsonPbCodec) Name() string { _ = "STUB: not implemented"; return "" }

func (c *jsonPbCodec) Close() error { _ = "STUB: not implemented"; return nil }

func getStreamingMode(fnSvc *dproto.MethodDescriptor) serviceinfo.StreamingMode {
	_ = "STUB: not implemented"
	return *new(serviceinfo.StreamingMode)
}
