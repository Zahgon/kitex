package descriptor

import (
	"context"
)

type ValueMapping interface {
	Request(ctx context.Context, val interface{}, field *FieldDescriptor) (interface{}, error)
	Response(ctx context.Context, val interface{}, field *FieldDescriptor) (interface{}, error)
}

type NewValueMapping func(value string) ValueMapping

var APIJSConvAnnotation = NewBAMAnnotation("api.js_conv", NewAPIJSConv)

type apiJSConv struct{}

var NewAPIJSConv NewValueMapping = func(value string) ValueMapping {

	return &apiJSConv{}
}

func (m *apiJSConv) Request(ctx context.Context, val interface{}, field *FieldDescriptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *apiJSConv) Response(ctx context.Context, val interface{}, field *FieldDescriptor) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
