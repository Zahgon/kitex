package bthrift

import (
	thrift "github.com/cloudwego/kitex/pkg/protocol/bthrift/apache"
)

var (
	Binary binaryProtocol
	_      BTProtocol = binaryProtocol{}
)

type binaryProtocol struct{}

func SetSpanCache(enable bool) { _ = "STUB: not implemented"; return }

func (binaryProtocol) WriteMessageBegin(buf []byte, name string, typeID thrift.TMessageType, seqid int32) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) WriteMessageEnd(buf []byte) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteStructBegin(buf []byte, name string) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) WriteStructEnd(buf []byte) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteFieldBegin(buf []byte, name string, typeID thrift.TType, id int16) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) WriteFieldEnd(buf []byte) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteFieldStop(buf []byte) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteMapBegin(buf []byte, keyType, valueType thrift.TType, size int) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) WriteMapEnd(buf []byte) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteListBegin(buf []byte, elemType thrift.TType, size int) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) WriteListEnd(buf []byte) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteSetBegin(buf []byte, elemType thrift.TType, size int) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) WriteSetEnd(buf []byte) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteBool(buf []byte, value bool) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteByte(buf []byte, value int8) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteI16(buf []byte, value int16) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteI32(buf []byte, value int32) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteI64(buf []byte, value int64) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteDouble(buf []byte, value float64) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) WriteString(buf []byte, value string) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) WriteBinary(buf, value []byte) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) WriteStringNocopy(buf []byte, binaryWriter BinaryWriter, value string) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) WriteBinaryNocopy(buf []byte, binaryWriter BinaryWriter, value []byte) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) MessageBeginLength(name string, _ thrift.TMessageType, _ int32) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) MessageEndLength() int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) StructBeginLength(name string) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) StructEndLength() int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) FieldBeginLength(name string, typeID thrift.TType, id int16) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) FieldEndLength() int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) FieldStopLength() int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) MapBeginLength(keyType, valueType thrift.TType, size int) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) MapEndLength() int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) ListBeginLength(elemType thrift.TType, size int) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) ListEndLength() int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) SetBeginLength(elemType thrift.TType, size int) int {
	_ = "STUB: not implemented"
	return 0
}

func (binaryProtocol) SetEndLength() int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) BoolLength(value bool) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) ByteLength(value int8) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) I16Length(value int16) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) I32Length(value int32) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) I64Length(value int64) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) DoubleLength(value float64) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) StringLength(value string) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) BinaryLength(value []byte) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) StringLengthNocopy(value string) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) BinaryLengthNocopy(value []byte) int { _ = "STUB: not implemented"; return 0 }

func (binaryProtocol) ReadMessageBegin(buf []byte) (name string, typeID thrift.TMessageType, seqid int32, length int, err error) {
	_ = "STUB: not implemented"
	return "", *new(thrift.TMessageType), 0, 0, nil
}

func (binaryProtocol) ReadMessageEnd(_ []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (binaryProtocol) ReadStructBegin(_ []byte) (name string, length int, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (binaryProtocol) ReadStructEnd(_ []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (binaryProtocol) ReadFieldBegin(buf []byte) (name string, typeID thrift.TType, id int16, length int, err error) {
	_ = "STUB: not implemented"
	return "", *new(thrift.TType), 0, 0, nil
}

func (binaryProtocol) ReadFieldEnd(_ []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (binaryProtocol) ReadMapBegin(buf []byte) (keyType, valueType thrift.TType, size, length int, err error) {
	_ = "STUB: not implemented"
	return *new(thrift.TType), *new(thrift.TType), 0, 0, nil
}

func (binaryProtocol) ReadMapEnd(_ []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (binaryProtocol) ReadListBegin(buf []byte) (elemType thrift.TType, size, length int, err error) {
	_ = "STUB: not implemented"
	return *new(thrift.TType), 0, 0, nil
}

func (binaryProtocol) ReadListEnd(_ []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (binaryProtocol) ReadSetBegin(buf []byte) (elemType thrift.TType, size, length int, err error) {
	_ = "STUB: not implemented"
	return *new(thrift.TType), 0, 0, nil
}

func (binaryProtocol) ReadSetEnd(_ []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (binaryProtocol) ReadBool(buf []byte) (value bool, length int, err error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func (binaryProtocol) ReadByte(buf []byte) (value int8, length int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (binaryProtocol) ReadI16(buf []byte) (value int16, length int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (binaryProtocol) ReadI32(buf []byte) (value int32, length int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (binaryProtocol) ReadI64(buf []byte) (value int64, length int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (binaryProtocol) ReadDouble(buf []byte) (value float64, length int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (binaryProtocol) ReadString(buf []byte) (value string, length int, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (binaryProtocol) ReadBinary(buf []byte) (value []byte, length int, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (binaryProtocol) Skip(buf []byte, fieldType thrift.TType) (length int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
