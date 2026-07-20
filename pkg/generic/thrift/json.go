package thrift

import (
	"context"

	"github.com/cloudwego/dynamicgo/conv"
	"github.com/cloudwego/dynamicgo/conv/j2t"
	dthrift "github.com/cloudwego/dynamicgo/thrift"
	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/protocol/thrift"
	"github.com/cloudwego/gopkg/protocol/thrift/base"
	"github.com/tidwall/gjson"

	"github.com/cloudwego/kitex/pkg/generic/descriptor"
)

type JSONReaderWriter struct {
	*ReadJSON
	*WriteJSON
}

func NewJsonReaderWriter(svc *descriptor.ServiceDescriptor) *JSONReaderWriter {
	_ = "STUB: not implemented"
	return nil
}

func NewWriteJSON(svc *descriptor.ServiceDescriptor) *WriteJSON {
	_ = "STUB: not implemented"
	return nil
}

const voidWholeLen = 5

var _ = wrapJSONWriter

type WriteJSON struct {
	svcDsc                 *descriptor.ServiceDescriptor
	base64Binary           bool
	convOpts               conv.Options
	convOptsWithThriftBase conv.Options
	dynamicgoEnabled       bool
}

var _ MessageWriter = (*WriteJSON)(nil)

func (m *WriteJSON) SetBase64Binary(enable bool) { _ = "STUB: not implemented"; return }

func (m *WriteJSON) SetDynamicGo(convOpts, convOptsWithThriftBase *conv.Options) {
	_ = "STUB: not implemented"
	return
}

func (m *WriteJSON) originalWrite(ctx context.Context, out bufiox.Writer, msg interface{}, method string, isClient bool, requestBase *base.Base) error {
	_ = "STUB: not implemented"
	return nil
}

func NewReadJSON(svc *descriptor.ServiceDescriptor) *ReadJSON {
	_ = "STUB: not implemented"
	return nil
}

type ReadJSON struct {
	svc                   *descriptor.ServiceDescriptor
	binaryWithBase64      bool
	convOpts              conv.Options
	convOptsWithException conv.Options
	dynamicgoEnabled      bool
}

var _ MessageReader = (*ReadJSON)(nil)

func (m *ReadJSON) SetBinaryWithBase64(enable bool) { _ = "STUB: not implemented"; return }

func (m *ReadJSON) SetDynamicGo(convOpts, convOptsWithException *conv.Options) {
	_ = "STUB: not implemented"
	return
}

func (m *ReadJSON) Read(ctx context.Context, method string, isClient bool, dataLen int, in bufiox.Reader) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *ReadJSON) originalRead(ctx context.Context, method string, isClient bool, in bufiox.Reader) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func removePrefixAndSuffix(buf []byte) []byte { _ = "STUB: not implemented"; return nil }

func jsonWriter(ctx context.Context, body *gjson.Result, typeDsc *descriptor.TypeDescriptor, opt *writerOption, bw *thrift.BufferWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func structReader(ctx context.Context, typeDesc *descriptor.TypeDescriptor, opt *readerOption, br *thrift.BufferReader) (v interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *WriteJSON) Write(ctx context.Context, out bufiox.Writer, msg interface{}, method string, isClient bool, requestBase *base.Base) error {
	_ = "STUB: not implemented"
	return nil
}

type MsgType int

func writeFields(ctx context.Context, out bufiox.Writer, dynamicgoTypeDsc *dthrift.TypeDescriptor, cv *j2t.BinaryConv, transBuff []byte, isClient bool) error {
	_ = "STUB: not implemented"
	return nil
}

func writeUnwrappedFields(ctx context.Context, out bufiox.Writer, dynamicgoTypeDsc *dthrift.TypeDescriptor, cv *j2t.BinaryConv, transBuff []byte) error {
	_ = "STUB: not implemented"
	return nil
}
