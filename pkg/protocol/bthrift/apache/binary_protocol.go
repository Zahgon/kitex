package apache

import (
	"context"
	"sync"

	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/protocol/thrift"
)

var (
	_ TProtocol = (*BinaryProtocol)(nil)

	bpPool = sync.Pool{
		New: func() interface{} {
			return &BinaryProtocol{}
		},
	}
)

type BinaryProtocol struct {
	r *thrift.BufferReader
	w *thrift.BufferWriter

	br bufiox.Reader
	bw bufiox.Writer
}

func NewBinaryProtocol(r bufiox.Reader, w bufiox.Writer) *BinaryProtocol {
	_ = "STUB: not implemented"
	return nil
}

func (p *BinaryProtocol) Recycle() { _ = "STUB: not implemented"; return }

func (p *BinaryProtocol) WriteMessageBegin(name string, typeID TMessageType, seqID int32) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *BinaryProtocol) WriteMessageEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteStructBegin(name string) error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteStructEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteFieldBegin(name string, typeID TType, id int16) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *BinaryProtocol) WriteFieldEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteFieldStop() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteMapBegin(keyType, valueType TType, size int) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *BinaryProtocol) WriteMapEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteListBegin(elemType TType, size int) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *BinaryProtocol) WriteListEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteSetBegin(elemType TType, size int) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *BinaryProtocol) WriteSetEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteBool(value bool) error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteByte(value int8) error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteI16(value int16) error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteI32(value int32) error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteI64(value int64) error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteDouble(value float64) error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteString(value string) error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) WriteBinary(value []byte) error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) ReadMessageBegin() (name string, typeID TMessageType, seqID int32, err error) {
	_ = "STUB: not implemented"
	return "", *new(TMessageType), 0, nil
}

func (p *BinaryProtocol) ReadMessageEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) ReadStructBegin() (name string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *BinaryProtocol) ReadStructEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) ReadFieldBegin() (name string, typeID TType, id int16, err error) {
	_ = "STUB: not implemented"
	return "", *new(TType), 0, nil
}

func (p *BinaryProtocol) ReadFieldEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) ReadMapBegin() (kType, vType TType, size int, err error) {
	_ = "STUB: not implemented"
	return *new(TType), *new(TType), 0, nil
}

func (p *BinaryProtocol) ReadMapEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) ReadListBegin() (elemType TType, size int, err error) {
	_ = "STUB: not implemented"
	return *new(TType), 0, nil
}

func (p *BinaryProtocol) ReadListEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) ReadSetBegin() (elemType TType, size int, err error) {
	_ = "STUB: not implemented"
	return *new(TType), 0, nil
}

func (p *BinaryProtocol) ReadSetEnd() error { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) ReadBool() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (p *BinaryProtocol) ReadByte() (value int8, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *BinaryProtocol) ReadI16() (value int16, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *BinaryProtocol) ReadI32() (value int32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *BinaryProtocol) ReadI64() (value int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *BinaryProtocol) ReadDouble() (value float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *BinaryProtocol) ReadString() (value string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p *BinaryProtocol) ReadBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *BinaryProtocol) Flush(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *BinaryProtocol) Skip(fieldType TType) (err error) { _ = "STUB: not implemented"; return nil }

func (p *BinaryProtocol) Transport() TTransport { _ = "STUB: not implemented"; return *new(TTransport) }

func (p *BinaryProtocol) GetBufioxReader() bufiox.Reader {
	_ = "STUB: not implemented"
	return *new(bufiox.Reader)
}

func (p *BinaryProtocol) GetBufioxWriter() bufiox.Writer {
	_ = "STUB: not implemented"
	return *new(bufiox.Writer)
}

type ttransportByteBuffer struct{}

func (ttransportByteBuffer) Close() error { _ = "STUB: not implemented"; return nil }
func (ttransportByteBuffer) Flush(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}
func (ttransportByteBuffer) IsOpen() bool             { _ = "STUB: not implemented"; return false }
func (ttransportByteBuffer) Open() error              { _ = "STUB: not implemented"; return nil }
func (p ttransportByteBuffer) RemainingBytes() uint64 { _ = "STUB: not implemented"; return 0 }
func (ttransportByteBuffer) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
func (ttransportByteBuffer) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
