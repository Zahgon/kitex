package descriptor

import (
	"context"
)

type HTTPMapping interface {
	Request(ctx context.Context, req *HTTPRequest, field *FieldDescriptor) (interface{}, bool, error)

	Response(ctx context.Context, resp *HTTPResponse, field *FieldDescriptor, val interface{}) error
}

type NewHTTPMapping func(value string) HTTPMapping

var DefaultNewMapping = NewAPIBody

var (
	APIQueryAnnotation = NewBAMAnnotation("api.query", NewAPIQuery)

	APIPathAnnotation = NewBAMAnnotation("api.path", NewAPIPath)

	APIHeaderAnnotation = NewBAMAnnotation("api.header", NewAPIHeader)

	APICookieAnnotation = NewBAMAnnotation("api.cookie", NewAPICookie)

	APIBodyAnnotation = NewBAMAnnotation("api.body", NewAPIBody)

	APIHttpCodeAnnotation = NewBAMAnnotation("api.http_code", NewAPIHTTPCode)

	APINoneAnnotation = NewBAMAnnotation("api.none", NewAPINone)

	APIRawBodyAnnotation = NewBAMAnnotation("api.raw_body", NewAPIRawBody)
)

type apiQuery struct {
	value string
}

var NewAPIQuery NewHTTPMapping = func(value string) HTTPMapping {
	return &apiQuery{value}
}

func (m *apiQuery) Request(ctx context.Context, req *HTTPRequest, field *FieldDescriptor) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (*apiQuery) Response(ctx context.Context, resp *HTTPResponse, field *FieldDescriptor, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type apiPath struct {
	value string
}

var NewAPIPath NewHTTPMapping = func(value string) HTTPMapping {
	return &apiPath{value}
}

func (m *apiPath) Request(ctx context.Context, req *HTTPRequest, field *FieldDescriptor) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (*apiPath) Response(ctx context.Context, resp *HTTPResponse, field *FieldDescriptor, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type apiHeader struct {
	value string
}

var NewAPIHeader NewHTTPMapping = func(value string) HTTPMapping {
	return &apiHeader{value}
}

func (m *apiHeader) Request(ctx context.Context, req *HTTPRequest, field *FieldDescriptor) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *apiHeader) Response(ctx context.Context, resp *HTTPResponse, field *FieldDescriptor, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type apiCookie struct {
	value string
}

var NewAPICookie NewHTTPMapping = func(value string) HTTPMapping {
	return &apiCookie{value}
}

func (m *apiCookie) Request(ctx context.Context, req *HTTPRequest, field *FieldDescriptor) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *apiCookie) Response(ctx context.Context, resp *HTTPResponse, field *FieldDescriptor, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type apiBody struct {
	value string
}

var NewAPIBody NewHTTPMapping = func(value string) HTTPMapping {
	return &apiBody{value}
}

func (m *apiBody) Request(ctx context.Context, req *HTTPRequest, field *FieldDescriptor) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *apiBody) Response(ctx context.Context, resp *HTTPResponse, field *FieldDescriptor, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type apiHTTPCode struct{}

var NewAPIHTTPCode NewHTTPMapping = func(value string) HTTPMapping {
	return &apiHTTPCode{}
}

func (m *apiHTTPCode) Request(ctx context.Context, req *HTTPRequest, field *FieldDescriptor) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *apiHTTPCode) Response(ctx context.Context, resp *HTTPResponse, field *FieldDescriptor, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type apiNone struct{}

var NewAPINone NewHTTPMapping = func(value string) HTTPMapping {
	return &apiNone{}
}

func (m *apiNone) Request(ctx context.Context, req *HTTPRequest, field *FieldDescriptor) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *apiNone) Response(ctx context.Context, resp *HTTPResponse, field *FieldDescriptor, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type apiRawBody struct{}

var NewAPIRawBody NewHTTPMapping = func(value string) HTTPMapping {
	return &apiRawBody{}
}

func (m *apiRawBody) Request(ctx context.Context, req *HTTPRequest, field *FieldDescriptor) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (m *apiRawBody) Response(ctx context.Context, resp *HTTPResponse, field *FieldDescriptor, val interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
