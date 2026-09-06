package test

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/client/genericclient"
	kt "github.com/cloudwego/kitex/internal/mocks/thrift"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/transport"
)

type Simple struct {
	ByteField   int8    `thrift:"ByteField,1" json:"ByteField"`
	I64Field    int64   `thrift:"I64Field,2" json:"I64Field"`
	DoubleField float64 `thrift:"DoubleField,3" json:"DoubleField"`
	I32Field    int32   `thrift:"I32Field,4" json:"I32Field"`
	StringField string  `thrift:"StringField,5" json:"StringField"`
	BinaryField []byte  `thrift:"BinaryField,6" json:"BinaryField"`
}

type Nesting struct {
	String_         string             `thrift:"String,1" json:"String"`
	ListSimple      []*Simple          `thrift:"ListSimple,2" json:"ListSimple"`
	Double          float64            `thrift:"Double,3" json:"Double"`
	I32             int32              `thrift:"I32,4" json:"I32"`
	ListI32         []int32            `thrift:"ListI32,5" json:"ListI32"`
	I64             int64              `thrift:"I64,6" json:"I64"`
	MapStringString map[string]string  `thrift:"MapStringString,7" json:"MapStringString"`
	SimpleStruct    *Simple            `thrift:"SimpleStruct,8" json:"SimpleStruct"`
	MapI32I64       map[int32]int64    `thrift:"MapI32I64,9" json:"MapI32I64"`
	ListString      []string           `thrift:"ListString,10" json:"ListString"`
	Binary          []byte             `thrift:"Binary,11" json:"Binary"`
	MapI64String    map[int64]string   `thrift:"MapI64String,12" json:"MapI64String"`
	ListI64         []int64            `thrift:"ListI64,13" json:"ListI64"`
	Byte            int8               `thrift:"Byte,14" json:"Byte"`
	MapStringSimple map[string]*Simple `thrift:"MapStringSimple,15" json:"MapStringSimple"`
}

func getString() string { _ = "STUB: not implemented"; return "" }

func getBytes() []byte { _ = "STUB: not implemented"; return nil }

func getSimpleValue() *Simple { _ = "STUB: not implemented"; return nil }

func getNestingValue() *Nesting { _ = "STUB: not implemented"; return nil }

func newGenericClient(tp transport.Protocol, destService string, g generic.Generic, targetIPPort string) genericclient.Client {
	_ = "STUB: not implemented"
	return *new(genericclient.Client)
}

func newGenericServer(g generic.Generic, addr net.Addr, handler generic.Service) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

type GenericServiceBinaryEchoImpl struct{}

const mockMyMsg = "my msg"

func (g *GenericServiceBinaryEchoImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GenericServiceBenchmarkImpl struct{}

func (g *GenericServiceBenchmarkImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GenericServiceAnnotationImpl struct{}

func (g *GenericServiceAnnotationImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func assertErr(field string, expected, actual interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	mockReq  = `{"Msg":"hello","strMap":{"mk1":"mv1","mk2":"mv2"},"strList":["lv1","lv2"]} `
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

func newMockExceptionTestArgs() interface{} { _ = "STUB: not implemented"; return nil }

func newMockExceptionTestResult() interface{} { _ = "STUB: not implemented"; return nil }

func exceptionHandler(ctx context.Context, handler, args, result interface{}) error {
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
