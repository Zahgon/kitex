package thrift

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/dynamicgo/conv"
	"github.com/cloudwego/dynamicgo/conv/t2j"
	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/protocol/thrift/base"

	"github.com/cloudwego/kitex/pkg/generic/descriptor"
)

type HTTPReaderWriter struct {
	*ReadHTTPResponse
	*WriteHTTPRequest
}

func NewHTTPReaderWriter(svc *descriptor.ServiceDescriptor) *HTTPReaderWriter {
	_ = "STUB: not implemented"
	return nil
}

type WriteHTTPRequest struct {
	svc                    *descriptor.ServiceDescriptor
	binaryWithBase64       bool
	convOpts               conv.Options
	convOptsWithThriftBase conv.Options
	dynamicgoEnabled       bool
}

var (
	_          MessageWriter = (*WriteHTTPRequest)(nil)
	customJson               = sonic.Config{
		EscapeHTML: true,
		UseNumber:  true,
		CopyString: true,
	}.Froze()
)

func NewWriteHTTPRequest(svc *descriptor.ServiceDescriptor) *WriteHTTPRequest {
	_ = "STUB: not implemented"
	return nil
}

func (w *WriteHTTPRequest) SetBinaryWithBase64(enable bool) { _ = "STUB: not implemented"; return }

func (w *WriteHTTPRequest) SetDynamicGo(convOpts, convOptsWithThriftBase *conv.Options) {
	_ = "STUB: not implemented"
	return
}

func (w *WriteHTTPRequest) originalWrite(ctx context.Context, out bufiox.Writer, msg interface{}, requestBase *base.Base) error {
	_ = "STUB: not implemented"
	return nil
}

type ReadHTTPResponse struct {
	svc                   *descriptor.ServiceDescriptor
	base64Binary          bool
	dynamicgoEnabled      bool
	useRawBodyForHTTPResp bool
	t2jBinaryConv         t2j.BinaryConv
}

var _ MessageReader = (*ReadHTTPResponse)(nil)

func NewReadHTTPResponse(svc *descriptor.ServiceDescriptor) *ReadHTTPResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReadHTTPResponse) SetBase64Binary(enable bool) { _ = "STUB: not implemented"; return }

func (r *ReadHTTPResponse) SetUseRawBodyForHTTPResp(useRawBodyForHTTPResp bool) {
	_ = "STUB: not implemented"
	return
}

func (r *ReadHTTPResponse) SetDynamicGo(convOpts *conv.Options) { _ = "STUB: not implemented"; return }

func (r *ReadHTTPResponse) Read(ctx context.Context, method string, isClient bool, dataLen int, in bufiox.Reader) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *ReadHTTPResponse) originalRead(ctx context.Context, method string, in bufiox.Reader) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *WriteHTTPRequest) Write(ctx context.Context, out bufiox.Writer, msg interface{}, method string, isClient bool, requestBase *base.Base) error {
	_ = "STUB: not implemented"
	return nil
}
