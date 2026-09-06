package thrift

import (
	"context"

	"github.com/cloudwego/gopkg/protocol/thrift"
	"github.com/cloudwego/gopkg/protocol/thrift/base"
	"github.com/tidwall/gjson"

	"github.com/cloudwego/kitex/internal/generic/proto"
	"github.com/cloudwego/kitex/pkg/generic/descriptor"
)

type writerOption struct {
	requestBase *base.Base

	binaryWithBase64 bool
}

type writer func(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error

type fieldGetter func(val interface{}, field *descriptor.FieldDescriptor) (interface{}, bool)

var mapGetter fieldGetter = func(val interface{}, field *descriptor.FieldDescriptor) (interface{}, bool) {
	st := val.(map[string]interface{})
	ret, ok := st[field.FieldName()]
	return ret, ok
}

var pbGetter fieldGetter = func(val interface{}, field *descriptor.FieldDescriptor) (interface{}, bool) {
	st := val.(proto.Message)
	ret, err := st.TryGetFieldByNumber(int(field.ID))
	return ret, err == nil
}

func typeOf(sample interface{}, t *descriptor.TypeDescriptor, opt *writerOption) (descriptor.Type, writer, error) {
	_ = "STUB: not implemented"
	return *new(descriptor.Type), *new(writer), nil
}

func typeJSONOf(data *gjson.Result, t *descriptor.TypeDescriptor, opt *writerOption) (v interface{}, w writer, err error) {
	_ = "STUB: not implemented"
	return nil, *new(writer), nil
}

func getWriterAndWrite(ctx context.Context, elem interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func nextWriter(sample interface{}, t *descriptor.TypeDescriptor, opt *writerOption) (writer, error) {
	_ = "STUB: not implemented"
	return *new(writer), nil
}

func nextJSONWriter(data *gjson.Result, t *descriptor.TypeDescriptor, opt *writerOption) (interface{}, writer, error) {
	_ = "STUB: not implemented"
	return nil, *new(writer), nil
}

func writeEmptyValue(out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func wrapStructWriter(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func structWriter(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func wrapJSONWriter(ctx context.Context, val *gjson.Result, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeVoid(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeBool(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeInt8(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeInt16(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeInt32(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeInt64(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeJSONNumber(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeJSONFloat64(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeFloat64(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeString(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeBase64Binary(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeBinary(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeBinaryList(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeList(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeJSONList(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeInterfaceMap(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeStringMap(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeStringJSONMap(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeRequestBase(ctx context.Context, val interface{}, out *thrift.BufferWriter, field *descriptor.FieldDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeStruct(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeHTTPRequest(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}

func writeJSON(ctx context.Context, val interface{}, out *thrift.BufferWriter, t *descriptor.TypeDescriptor, opt *writerOption) error {
	_ = "STUB: not implemented"
	return nil
}
