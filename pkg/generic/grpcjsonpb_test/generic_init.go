package test

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/internal/mocks/proto/kitex_gen/pbapi/mock"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/server"
)

var startChan chan struct{}

func init() {
	startChan = make(chan struct{})
	server.RegisterStartHook(func() {
		startChan <- struct{}{}
	})
}

func newGenericClient(g generic.Generic, targetIPPort string, cliOpts ...client.Option) genericclient.Client {
	_ = "STUB: not implemented"
	return *new(genericclient.Client)
}

func newMockTestServer(handler mock.Mock, addr net.Addr, opts ...server.Option) server.Server {
	_ = "STUB: not implemented"
	return *new(server.Server)
}

var _ mock.Mock = &StreamingTestImpl{}

type StreamingTestImpl struct{}

func (s *StreamingTestImpl) UnaryTest(ctx context.Context, req *mock.MockReq) (resp *mock.MockResp, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *StreamingTestImpl) ClientStreamingTest(stream mock.Mock_ClientStreamingTestServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *StreamingTestImpl) ServerStreamingTest(req *mock.MockReq, stream mock.Mock_ServerStreamingTestServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *StreamingTestImpl) BidirectionalStreamingTest(stream mock.Mock_BidirectionalStreamingTestServer) (err error) {
	_ = "STUB: not implemented"
	return nil
}
