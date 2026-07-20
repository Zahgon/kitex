package serviceinfo

import (
	"context"
)

type PayloadCodec int

const (
	Thrift PayloadCodec = iota
	Protobuf
	Hessian2
)

const (
	GenericService = "$GenericService"

	GenericMethod = "$GenericCall"

	PackageName = "PackageName"

	CombineServiceKey = "combine_service"

	CombineServiceName = "CombineService"
)

type GenericMethodFunc func(ctx context.Context, methodName string) MethodInfo

type ServiceInfo struct {
	PackageName string

	ServiceName string

	HandlerType interface{}

	Methods map[string]MethodInfo

	PayloadCodec PayloadCodec

	KiteXGenVersion string

	Extra map[string]interface{}

	GenericMethod GenericMethodFunc
}

func (i *ServiceInfo) GetPackageName() (pkg string) { _ = "STUB: not implemented"; return "" }

func (i *ServiceInfo) MethodInfo(ctx context.Context, name string) MethodInfo {
	_ = "STUB: not implemented"
	return *new(MethodInfo)
}

type StreamingMode int

const (
	StreamingNone          StreamingMode = 0b0000
	StreamingUnary         StreamingMode = 0b0001
	StreamingClient        StreamingMode = 0b0010
	StreamingServer        StreamingMode = 0b0100
	StreamingBidirectional StreamingMode = 0b0110
)

type MethodInfo interface {
	Handler() MethodHandler
	NewArgs() interface{}
	NewResult() interface{}
	OneWay() bool
	IsStreaming() bool
	StreamingMode() StreamingMode
}

type MethodHandler func(ctx context.Context, handler, args, result interface{}) error

type MethodInfoOption func(*methodInfo)

func WithStreamingMode(mode StreamingMode) MethodInfoOption {
	_ = "STUB: not implemented"
	return *new(MethodInfoOption)
}

func NewMethodInfo(methodHandler MethodHandler, newArgsFunc, newResultFunc func() interface{}, oneWay bool, opts ...MethodInfoOption) MethodInfo {
	_ = "STUB: not implemented"
	return *new(MethodInfo)
}

type methodInfo struct {
	handler       MethodHandler
	newArgsFunc   func() interface{}
	newResultFunc func() interface{}
	oneWay        bool
	isStreaming   bool
	streamingMode StreamingMode
}

func (m methodInfo) Handler() MethodHandler { _ = "STUB: not implemented"; return *new(MethodHandler) }

func (m methodInfo) NewArgs() interface{} { _ = "STUB: not implemented"; return nil }

func (m methodInfo) NewResult() interface{} { _ = "STUB: not implemented"; return nil }

func (m methodInfo) OneWay() bool { _ = "STUB: not implemented"; return false }

func (m methodInfo) IsStreaming() bool { _ = "STUB: not implemented"; return false }

func (m methodInfo) StreamingMode() StreamingMode {
	_ = "STUB: not implemented"
	return *new(StreamingMode)
}

func (p PayloadCodec) String() string { _ = "STUB: not implemented"; return "" }
