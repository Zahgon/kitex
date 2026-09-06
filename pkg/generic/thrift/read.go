package thrift

import (
	"context"
	"reflect"

	"github.com/cloudwego/gopkg/protocol/thrift"
	"github.com/jhump/protoreflect/desc"

	"github.com/cloudwego/kitex/internal/generic/proto"
	"github.com/cloudwego/kitex/pkg/generic/descriptor"
)

var emptyPbDsc = &desc.MessageDescriptor{}

type readerOption struct {
	forJSON bool

	throwException bool

	http                bool
	binaryWithBase64    bool
	binaryWithByteSlice bool

	setFieldsForEmptyStruct uint8

	pbDsc proto.MessageDescriptor
}

type reader func(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error)

type fieldSetter func(field *descriptor.FieldDescriptor, val interface{}) error

func getMapFieldSetter(st map[string]interface{}) fieldSetter {
	_ = "STUB: not implemented"
	return *new(fieldSetter)
}

func getPbFieldSetter(st proto.Message) fieldSetter {
	_ = "STUB: not implemented"
	return *new(fieldSetter)
}

func nextReader(tt descriptor.Type, t *descriptor.TypeDescriptor, opt *readerOption) (reader, error) {
	_ = "STUB: not implemented"
	return *new(reader), nil
}

func skipStructReader(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readVoid(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readDouble(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readBool(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readByte(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readInt16(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readInt32(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readInt64(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readString(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readBinary(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readBase64Binary(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readList(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readMap(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readInterfaceMap(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readStringMap(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readStruct(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readHTTPResponse(ctx context.Context, in *thrift.BufferReader, t *descriptor.TypeDescriptor, opt *readerOption) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unnestPb(opt *readerOption, fieldId int32) func() { _ = "STUB: not implemented"; return nil }

func readEmptyValue(t *descriptor.TypeDescriptor) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func getRType(t *descriptor.TypeDescriptor) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}
