package generic

import (
	"context"
	"errors"

	igeneric "github.com/cloudwego/kitex/internal/generic"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

var (
	errGenericCallNotImplemented     = errors.New("generic.ServiceV2 GenericCall not implemented")
	errClientStreamingNotImplemented = errors.New("generic.ServiceV2 ClientStreaming not implemented")
	errServerStreamingNotImplemented = errors.New("generic.ServiceV2 ServerStreaming not implemented")
	errBidiStreamingNotImplemented   = errors.New("generic.ServiceV2 BidiStreaming not implemented")
)

type Service interface {
	GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error)
}

type ServiceV2 struct {
	GenericCall func(ctx context.Context, service, method string, request interface{}) (response interface{}, err error)

	ClientStreaming func(ctx context.Context, service, method string, stream ClientStreamingServer) (err error)

	ServerStreaming func(ctx context.Context, service, method string, request interface{}, stream ServerStreamingServer) (err error)

	BidiStreaming func(ctx context.Context, service, method string, stream BidiStreamingServer) (err error)
}

func ServiceInfoWithGeneric(g Generic) *serviceinfo.ServiceInfo {
	_ = "STUB: not implemented"
	return nil
}

func callHandler(ctx context.Context, handler, arg, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func newGenericServiceCallArgs() interface{} { _ = "STUB: not implemented"; return nil }

func newGenericServiceCallResult() interface{} { _ = "STUB: not implemented"; return nil }

func clientStreamingHandlerGetter(mi serviceinfo.MethodInfo) serviceinfo.MethodHandler {
	_ = "STUB: not implemented"
	return *new(serviceinfo.MethodHandler)
}

func serverStreamingHandlerGetter(mi serviceinfo.MethodInfo) serviceinfo.MethodHandler {
	_ = "STUB: not implemented"
	return *new(serviceinfo.MethodHandler)
}

func bidiStreamingHandlerGetter(mi serviceinfo.MethodInfo) serviceinfo.MethodHandler {
	_ = "STUB: not implemented"
	return *new(serviceinfo.MethodHandler)
}

type WithCodec interface {
	SetCodec(codec interface{})
}

type Args = igeneric.Args

type Result = igeneric.Result

var (
	_ WithCodec = (*Args)(nil)
	_ WithCodec = (*Result)(nil)
)
