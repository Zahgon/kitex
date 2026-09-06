package test

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server"
)

func newGenericClient(destService string, g generic.Generic, targetIPPort string) genericclient.Client {
	_ = "STUB: not implemented"
	return *new(genericclient.Client)
}

func newGenericServer(g generic.Generic, addr net.Addr, handler generic.Service) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

type TestEchoService struct{}

func (g *TestEchoService) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestExampleMethodService struct{}

func (g *TestExampleMethodService) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestVoidService struct{}

func (g *TestVoidService) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestExampleMethod2Service struct{}

func (g *TestExampleMethod2Service) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestInt2FloatMethodService struct{}

type ExampleInt2Float struct {
	Int32   int32
	Float64 float64
	String_ string
	Int64   int64
	Subfix  float64
}

func (g *TestInt2FloatMethodService) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TestInt2FloatMethod2Service struct{}

func (g *TestInt2FloatMethod2Service) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getEchoReq() string { _ = "STUB: not implemented"; return "" }

func getBizErrReq() string { _ = "STUB: not implemented"; return "" }

func getEchoRes() string { _ = "STUB: not implemented"; return "" }

func getExampleMethodReq() string { _ = "STUB: not implemented"; return "" }

func getExampleMethodRes() string { _ = "STUB: not implemented"; return "" }

func getVoidReq() string { _ = "STUB: not implemented"; return "" }

func getVoidRes() string { _ = "STUB: not implemented"; return "" }

func getExampleMethod2Req() string { _ = "STUB: not implemented"; return "" }

func getExampleMethod2Res() string { _ = "STUB: not implemented"; return "" }

func getInt2FloatMethodReq() string { _ = "STUB: not implemented"; return "" }

func getInt2FloatMethodRes() string { _ = "STUB: not implemented"; return "" }

func getInt2FloatMethod2Res() string { _ = "STUB: not implemented"; return "" }

func getInt2FloatMethod2Req() string { _ = "STUB: not implemented"; return "" }
