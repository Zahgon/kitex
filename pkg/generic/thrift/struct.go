package thrift

import (
	"context"

	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/protocol/thrift/base"

	"github.com/cloudwego/kitex/pkg/generic/descriptor"
)

type StructReaderWriter struct {
	*ReadStruct
	*WriteStruct
}

func NewStructReaderWriter(svc *descriptor.ServiceDescriptor) *StructReaderWriter {
	_ = "STUB: not implemented"
	return nil
}

func NewStructReaderWriterForJSON(svc *descriptor.ServiceDescriptor) *StructReaderWriter {
	_ = "STUB: not implemented"
	return nil
}

func NewWriteStruct(svc *descriptor.ServiceDescriptor) *WriteStruct {
	_ = "STUB: not implemented"
	return nil
}

type WriteStruct struct {
	svcDsc           *descriptor.ServiceDescriptor
	binaryWithBase64 bool
}

var _ MessageWriter = (*WriteStruct)(nil)

func (m *WriteStruct) SetBinaryWithBase64(enable bool) { _ = "STUB: not implemented"; return }

func (m *WriteStruct) Write(ctx context.Context, out bufiox.Writer, msg interface{}, method string, isClient bool, requestBase *base.Base) error {
	_ = "STUB: not implemented"
	return nil
}

func NewReadStruct(svc *descriptor.ServiceDescriptor) *ReadStruct {
	_ = "STUB: not implemented"
	return nil
}

func NewReadStructForJSON(svc *descriptor.ServiceDescriptor) *ReadStruct {
	_ = "STUB: not implemented"
	return nil
}

type ReadStruct struct {
	svc                     *descriptor.ServiceDescriptor
	forJSON                 bool
	binaryWithBase64        bool
	binaryWithByteSlice     bool
	setFieldsForEmptyStruct uint8
}

var _ MessageReader = (*ReadStruct)(nil)

func (m *ReadStruct) SetBinaryOption(base64, byteSlice bool) { _ = "STUB: not implemented"; return }

func (m *ReadStruct) SetSetFieldsForEmptyStruct(mode uint8) { _ = "STUB: not implemented"; return }

func (m *ReadStruct) Read(ctx context.Context, method string, isClient bool, dataLen int, in bufiox.Reader) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
