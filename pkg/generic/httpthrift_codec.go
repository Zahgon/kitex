package generic

import (
	"net/http"
	"sync/atomic"

	"github.com/cloudwego/dynamicgo/conv"

	"github.com/cloudwego/kitex/pkg/generic/descriptor"
	"github.com/cloudwego/kitex/pkg/generic/thrift"
)

var _ Closer = &httpThriftCodec{}

type HTTPRequest = descriptor.HTTPRequest

type HTTPResponse = descriptor.HTTPResponse

type httpThriftCodec struct {
	svcDsc                 atomic.Value
	provider               DescriptorProvider
	binaryWithBase64       bool
	convOpts               conv.Options
	convOptsWithThriftBase conv.Options
	dynamicgoEnabled       bool
	useRawBodyForHTTPResp  bool
	svcName                atomic.Value
	combineService         atomic.Value
	readerWriter           atomic.Value
}

func newHTTPThriftCodec(p DescriptorProvider, opts *Options) *httpThriftCodec {
	_ = "STUB: not implemented"
	return nil
}

func (c *httpThriftCodec) update() { _ = "STUB: not implemented"; return }

func (c *httpThriftCodec) updateMessageReaderWriter() (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *httpThriftCodec) configureMessageReaderWriter(svc *descriptor.ServiceDescriptor) {
	_ = "STUB: not implemented"
	return
}

func (c *httpThriftCodec) getMessageReaderWriter() interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *httpThriftCodec) configureHTTPRequestWriter(writer *thrift.WriteHTTPRequest) {
	_ = "STUB: not implemented"
	return
}

func (c *httpThriftCodec) configureHTTPResponseReader(reader *thrift.ReadHTTPResponse) {
	_ = "STUB: not implemented"
	return
}

func (c *httpThriftCodec) getMethodByReq(req interface{}) (methodName string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *httpThriftCodec) getMethod(name string) (Method, error) {
	_ = "STUB: not implemented"
	return *new(Method), nil
}

func (c *httpThriftCodec) Name() string { _ = "STUB: not implemented"; return "" }

func (c *httpThriftCodec) Close() error { _ = "STUB: not implemented"; return nil }

func FromHTTPRequest(req *http.Request) (*HTTPRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
