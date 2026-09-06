package proto

import (
	"context"

	"github.com/cloudwego/dynamicgo/conv"
	dproto "github.com/cloudwego/dynamicgo/proto"
)

type JSONReaderWriter struct {
	*ReadJSON
	*WriteJSON
}

func NewJsonReaderWriter(svc *dproto.ServiceDescriptor, convOpts *conv.Options) *JSONReaderWriter {
	_ = "STUB: not implemented"
	return nil
}

func NewWriteJSON(svc *dproto.ServiceDescriptor, convOpts *conv.Options) *WriteJSON {
	_ = "STUB: not implemented"
	return nil
}

type WriteJSON struct {
	svcDsc            *dproto.ServiceDescriptor
	dynamicgoConvOpts *conv.Options
}

var _ MessageWriter = (*WriteJSON)(nil)

func (m *WriteJSON) Write(ctx context.Context, msg interface{}, method string, isClient bool) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewReadJSON(svc *dproto.ServiceDescriptor, convOpts *conv.Options) *ReadJSON {
	_ = "STUB: not implemented"
	return nil
}

type ReadJSON struct {
	dynamicgoConvOpts *conv.Options
	dynamicgoSvcDsc   *dproto.ServiceDescriptor
}

var _ MessageReader = (*ReadJSON)(nil)

func (m *ReadJSON) Read(ctx context.Context, method string, isClient bool, actualMsgBuf []byte) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
