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

var reqMsg = map[string]interface{}{
	"Msg": "hello",
	"InnerBase": map[string]interface{}{
		"Base": map[string]interface{}{
			"LogID": "log_id_inner",
		},
	},
	"Base": map[string]interface{}{
		"LogID": "log_id",
	},
}

var errResp = "Test Error"

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

type GenericServiceWithBase64Binary struct{}

func (g *GenericServiceWithBase64Binary) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GenericServiceWithByteSliceImpl struct{}

func (g *GenericServiceWithByteSliceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GenericServiceErrorImpl struct{}

func (g *GenericServiceErrorImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GenericServicePingImpl struct{}

func (g *GenericServicePingImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GenericServiceOnewayImpl struct{}

func (g *GenericServiceOnewayImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GenericServiceVoidImpl struct{}

func (g *GenericServiceVoidImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	mockReq = map[string]interface{}{
		"Msg": "hello",
		"strMap": map[interface{}]interface{}{
			"mk1": "mv1",
			"mk2": "mv2",
		},
		"strList": []interface{}{
			"lv1", "lv2",
		},
	}
	mockResp = "this is response"
)

func newMockServer(handler kt.Mock, addr net.Addr, opts ...server.Option) server.Server {
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

type mockImpl struct{}

func (m *mockImpl) Test(ctx context.Context, req *kt.MockReq) (r string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (m *mockImpl) ExceptionTest(ctx context.Context, req *kt.MockReq) (r string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}
