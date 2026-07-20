package generic

import (
	"net/http"
	"sync/atomic"
)

var _ Closer = &httpPbThriftCodec{}

type httpPbThriftCodec struct {
	svcDsc         atomic.Value
	pbSvcDsc       atomic.Value
	provider       DescriptorProvider
	pbProvider     PbDescriptorProvider
	svcName        atomic.Value
	combineService atomic.Value
	readerWriter   atomic.Value
}

func newHTTPPbThriftCodec(p DescriptorProvider, pbp PbDescriptorProvider) *httpPbThriftCodec {
	_ = "STUB: not implemented"
	return nil
}

func (c *httpPbThriftCodec) update() { _ = "STUB: not implemented"; return }

func (c *httpPbThriftCodec) getMethodByReq(req interface{}) (methodName string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *httpPbThriftCodec) getMethod(name string) (Method, error) {
	_ = "STUB: not implemented"
	return *new(Method), nil
}

func (c *httpPbThriftCodec) getMessageReaderWriter() interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (c *httpPbThriftCodec) Name() string { _ = "STUB: not implemented"; return "" }

func (c *httpPbThriftCodec) Close() error { _ = "STUB: not implemented"; return nil }

func FromHTTPPbRequest(req *http.Request) (*HTTPRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
