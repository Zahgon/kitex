package test

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	kt "github.com/cloudwego/kitex/internal/mocks/thrift"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/server"
)

var (
	reqMsg  = "Hello Kitex"
	respMsg = "Hi, I am Kitex"
	errResp = "Test Error"
)

func newGenericClient(destService string, g generic.Generic, targetIPPort string, opts ...client.Option) genericclient.Client {
	_ = "STUB: not implemented"
	return *new(genericclient.Client)
}

func newGenericServer(g generic.Generic, addr net.Addr, handler generic.Service) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

type GenericServiceImpl struct{}

func (g *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GenericServiceErrorImpl struct{}

func (g *GenericServiceErrorImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GenericServiceBizErrorImpl struct{}

func (g *GenericServiceBizErrorImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GenericServiceMockImpl struct{}

func (g *GenericServiceMockImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewMockServer(handler kt.Mock, addr net.Addr, opts ...server.Option) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

func serviceInfo() *serviceinfo.ServiceInfo { _ = "STUB: not implemented"; return nil }

func newMockTestArgs() interface{} { _ = "STUB: not implemented"; return nil }

func newMockTestResult() interface{} { _ = "STUB: not implemented"; return nil }

func testHandler(ctx context.Context, handler, arg, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type MockImpl struct{}

func (m *MockImpl) Test(ctx context.Context, req *kt.MockReq) (r string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *MockImpl) ExceptionTest(ctx context.Context, req *kt.MockReq) (r string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func genBinaryResp(method string) []byte { _ = "STUB: not implemented"; return nil }
