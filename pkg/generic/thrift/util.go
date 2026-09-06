package thrift

import (
	"context"

	"github.com/cloudwego/kitex/pkg/generic/descriptor"
)

func assertType(expected, but descriptor.Type) error { _ = "STUB: not implemented"; return nil }

func splitType(t string) (pkg, name string) { _ = "STUB: not implemented"; return "", "" }

func requestMappingValue(ctx context.Context, req *descriptor.HTTPRequest, field *descriptor.FieldDescriptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildinTypeFromString(s string, t *descriptor.TypeDescriptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func buildinTypeIntoString(val interface{}) string { _ = "STUB: not implemented"; return "" }
