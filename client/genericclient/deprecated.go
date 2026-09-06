package genericclient

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

func StreamingServiceInfo(g generic.Generic) *serviceinfo.ServiceInfo {
	_ = "STUB: not implemented"
	return nil
}

type ClientStreaming interface {
	streaming.Stream
	Send(req interface{}) error
	CloseAndRecv() (resp interface{}, err error)
}

type ServerStreaming interface {
	streaming.Stream
	Recv() (resp interface{}, err error)
}

type BidirectionalStreaming interface {
	streaming.Stream
	Send(req interface{}) error
	Recv() (resp interface{}, err error)
}

func NewStreamingClient(destService string, g generic.Generic, opts ...client.Option) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func NewStreamingClientWithServiceInfo(destService string, g generic.Generic, svcInfo *serviceinfo.ServiceInfo, opts ...client.Option) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

type deprecatedClientStreamingClient struct {
	streaming.Stream
	method     string
	methodInfo serviceinfo.MethodInfo
}

func NewClientStreaming(ctx context.Context, genericCli Client, method string, callOpts ...callopt.Option) (ClientStreaming, error) {
	_ = "STUB: not implemented"
	return *new(ClientStreaming), nil
}

func (cs *deprecatedClientStreamingClient) Send(req interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *deprecatedClientStreamingClient) CloseAndRecv() (resp interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type deprecatedServerStreamingClient struct {
	streaming.Stream
	methodInfo serviceinfo.MethodInfo
}

func NewServerStreaming(ctx context.Context, genericCli Client, method string, req interface{}, callOpts ...callopt.Option) (ServerStreaming, error) {
	_ = "STUB: not implemented"
	return *new(ServerStreaming), nil
}

func (ss *deprecatedServerStreamingClient) Recv() (resp interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type deprecatedBidirectionalStreamingClient struct {
	streaming.Stream
	method     string
	methodInfo serviceinfo.MethodInfo
}

func NewBidirectionalStreaming(ctx context.Context, genericCli Client, method string, callOpts ...callopt.Option) (BidirectionalStreaming, error) {
	_ = "STUB: not implemented"
	return *new(BidirectionalStreaming), nil
}

func (bs *deprecatedBidirectionalStreamingClient) Send(req interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (bs *deprecatedBidirectionalStreamingClient) Recv() (resp interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getStream(ctx context.Context, genericCli *genericServiceClient, method string, callOpts ...callopt.Option) (streaming.Stream, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Stream), nil
}
