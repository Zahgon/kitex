package bthrift

import (
	"github.com/cloudwego/gopkg/protocol/thrift/unknownfields"
	"github.com/cloudwego/thriftgo/generator/golang/extension/unknown"
)

type UnknownField struct {
	Name    string
	ID      int16
	Type    int
	KeyType int
	ValType int
	Value   interface{}
}

func fromGopkgUnknownFields(ff []unknownfields.UnknownField) []UnknownField {
	_ = "STUB: not implemented"
	return nil
}

func toGopkgUnknownFields(ff []UnknownField) []unknownfields.UnknownField {
	_ = "STUB: not implemented"
	return nil
}

func GetUnknownFields(v interface{}) (fields []UnknownField, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ConvertUnknownFields(buf unknown.Fields) (fields []UnknownField, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UnknownFieldsLength(fs []UnknownField) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func WriteUnknownFields(buf []byte, fs []UnknownField) (offset int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
