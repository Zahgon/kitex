package thrift

import (
	"context"

	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/protocol/thrift/base"
	"github.com/jhump/protoreflect/desc"

	"github.com/cloudwego/kitex/pkg/generic/descriptor"
	"github.com/cloudwego/kitex/pkg/generic/proto"
)

type HTTPPbReaderWriter struct {
	*ReadHTTPPbResponse
	*WriteHTTPPbRequest
}

func NewHTTPPbReaderWriter(svc *descriptor.ServiceDescriptor, pbsvc proto.ServiceDescriptor) *HTTPPbReaderWriter {
	_ = "STUB: not implemented"
	return nil
}

type WriteHTTPPbRequest struct {
	svc   *descriptor.ServiceDescriptor
	pbSvc *desc.ServiceDescriptor
}

var _ MessageWriter = (*WriteHTTPPbRequest)(nil)

func NewWriteHTTPPbRequest(svc *descriptor.ServiceDescriptor, pbSvc *desc.ServiceDescriptor) *WriteHTTPPbRequest {
	_ = "STUB: not implemented"
	return nil
}

func (w *WriteHTTPPbRequest) Write(ctx context.Context, out bufiox.Writer, msg interface{}, method string, isClient bool, requestBase *base.Base) error {
	_ = "STUB: not implemented"
	return nil
}

type ReadHTTPPbResponse struct {
	svc   *descriptor.ServiceDescriptor
	pbSvc proto.ServiceDescriptor
}

var _ MessageReader = (*ReadHTTPPbResponse)(nil)

func NewReadHTTPPbResponse(svc *descriptor.ServiceDescriptor, pbSvc proto.ServiceDescriptor) *ReadHTTPPbResponse {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReadHTTPPbResponse) Read(ctx context.Context, method string, isClient bool, dataLen int, in bufiox.Reader) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
