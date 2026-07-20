package test

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/client/genericclient"
	kt "github.com/cloudwego/kitex/internal/mocks/thrift"
	"github.com/cloudwego/kitex/pkg/generic"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/cloudwego/kitex/server"
)

func newGenericStreamingClient(g generic.Generic, targetIPPort string) genericclient.Client {
	_ = "STUB: not implemented"
	return *new(genericclient.Client)
}

func newGenericClient(g generic.Generic, targetIPPort string) genericclient.Client {
	_ = "STUB: not implemented"
	return *new(genericclient.Client)
}

var _ kt.TestService = &StreamingTestImpl{}

type StreamingTestImpl struct{}

func (s *StreamingTestImpl) Echo(stream kt.TestService_EchoServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *StreamingTestImpl) EchoClient(stream kt.TestService_EchoClientServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *StreamingTestImpl) EchoServer(req *kt.Request, stream kt.TestService_EchoServerServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *StreamingTestImpl) EchoUnary(ctx context.Context, req *kt.Request) (resp *kt.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StreamingTestImpl) EchoPingPong(ctx context.Context, req *kt.Request) (resp *kt.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StreamingTestImpl) EchoBizException(stream kt.TestService_EchoBizExceptionServer) error {
	_ = "STUB: not implemented"
	return nil
}

func newMockServer(handler kt.TestService, addr net.Addr, opts ...server.Option) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

var serviceMethods = map[string]kitex.MethodInfo{
	"Echo": kitex.NewMethodInfo(
		echoHandler,
		newTestServiceEchoArgs,
		newTestServiceEchoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingBidirectional),
	),
	"EchoClient": kitex.NewMethodInfo(
		echoClientHandler,
		newTestServiceEchoClientArgs,
		newTestServiceEchoClientResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingClient),
	),
	"EchoServer": kitex.NewMethodInfo(
		echoServerHandler,
		newTestServiceEchoServerArgs,
		newTestServiceEchoServerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingServer),
	),
	"EchoUnary": kitex.NewMethodInfo(
		echoUnaryHandler,
		newTestServiceEchoUnaryArgs,
		newTestServiceEchoUnaryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"EchoBizException": kitex.NewMethodInfo(
		echoBizExceptionHandler,
		newTestServiceEchoBizExceptionArgs,
		newTestServiceEchoBizExceptionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingClient),
	),
	"EchoPingPong": kitex.NewMethodInfo(
		echoPingPongHandler,
		newTestServiceEchoPingPongArgs,
		newTestServiceEchoPingPongResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

func serviceInfo() *kitex.ServiceInfo { _ = "STUB: not implemented"; return nil }

func echoHandler(ctx context.Context, handler, arg, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type testServiceEchoServer struct {
	streaming.Stream
}

func (x *testServiceEchoServer) Send(m *kt.Response) error { _ = "STUB: not implemented"; return nil }

func (x *testServiceEchoServer) Recv() (*kt.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTestServiceEchoArgs() interface{} { _ = "STUB: not implemented"; return nil }

func newTestServiceEchoResult() interface{} { _ = "STUB: not implemented"; return nil }

func echoClientHandler(ctx context.Context, handler, arg, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type testServiceEchoClientServer struct {
	streaming.Stream
}

func (x *testServiceEchoClientServer) SendAndClose(m *kt.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *testServiceEchoClientServer) Recv() (*kt.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTestServiceEchoClientArgs() interface{} { _ = "STUB: not implemented"; return nil }

func newTestServiceEchoClientResult() interface{} { _ = "STUB: not implemented"; return nil }

func echoServerHandler(ctx context.Context, handler, arg, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type testServiceEchoServerServer struct {
	streaming.Stream
}

func (x *testServiceEchoServerServer) Send(m *kt.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func newTestServiceEchoServerArgs() interface{} { _ = "STUB: not implemented"; return nil }

func newTestServiceEchoServerResult() interface{} { _ = "STUB: not implemented"; return nil }

func echoUnaryHandler(ctx context.Context, handler, arg, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func newTestServiceEchoUnaryArgs() interface{} { _ = "STUB: not implemented"; return nil }

func newTestServiceEchoUnaryResult() interface{} { _ = "STUB: not implemented"; return nil }

func echoBizExceptionHandler(ctx context.Context, handler, arg, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type testServiceEchoBizExceptionServer struct {
	streaming.Stream
}

func (x *testServiceEchoBizExceptionServer) SendAndClose(m *kt.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *testServiceEchoBizExceptionServer) Recv() (*kt.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTestServiceEchoBizExceptionArgs() interface{} { _ = "STUB: not implemented"; return nil }

func newTestServiceEchoBizExceptionResult() interface{} { _ = "STUB: not implemented"; return nil }

func echoPingPongHandler(ctx context.Context, handler, arg, result interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func newTestServiceEchoPingPongArgs() interface{} { _ = "STUB: not implemented"; return nil }

func newTestServiceEchoPingPongResult() interface{} { _ = "STUB: not implemented"; return nil }
