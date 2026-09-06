package genericclient

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/client/callopt/streamcall"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

var _ Client = &genericServiceClient{}

func NewClient(destService string, g generic.Generic, opts ...client.Option) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func NewClientWithServiceInfo(destService string, g generic.Generic, svcInfo *serviceinfo.ServiceInfo, opts ...client.Option) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

type Client interface {
	generic.Closer

	GenericCall(ctx context.Context, method string, request interface{}, callOptions ...callopt.Option) (response interface{}, err error)

	ClientStreaming(ctx context.Context, method string, callOptions ...streamcall.Option) (ClientStreamingClient, error)

	ServerStreaming(ctx context.Context, method string, req interface{}, callOptions ...streamcall.Option) (ServerStreamingClient, error)

	BidirectionalStreaming(ctx context.Context, method string, callOptions ...streamcall.Option) (BidiStreamingClient, error)
}

type genericServiceClient struct {
	svcInfo *serviceinfo.ServiceInfo
	kClient client.Client
	sClient client.Streaming
	g       generic.Generic

	isBinaryGeneric bool

	getMethodFunc generic.GetMethodNameByRequestFunc
}

func (gc *genericServiceClient) GenericCall(ctx context.Context, method string, request interface{}, callOptions ...callopt.Option) (response interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gc *genericServiceClient) Close() error { _ = "STUB: not implemented"; return nil }

func (gc *genericServiceClient) ClientStreaming(ctx context.Context, method string, callOptions ...streamcall.Option) (ClientStreamingClient, error) {
	_ = "STUB: not implemented"
	return *new(ClientStreamingClient), nil
}

func (gc *genericServiceClient) ServerStreaming(ctx context.Context, method string, req interface{}, callOptions ...streamcall.Option) (ServerStreamingClient, error) {
	_ = "STUB: not implemented"
	return *new(ServerStreamingClient), nil
}

func (gc *genericServiceClient) BidirectionalStreaming(ctx context.Context, method string, callOptions ...streamcall.Option) (BidiStreamingClient, error) {
	_ = "STUB: not implemented"
	return *new(BidiStreamingClient), nil
}

type ClientStreamingClient interface {
	Send(ctx context.Context, req interface{}) error

	CloseAndRecv(ctx context.Context) (interface{}, error)

	Header() (streaming.Header, error)

	Trailer() (streaming.Trailer, error)

	CloseSend(ctx context.Context) error

	Context() context.Context

	Streaming() streaming.ClientStream
}

type clientStreamingClient struct {
	methodInfo serviceinfo.MethodInfo
	method     string
	streaming  streaming.ClientStream
}

func newClientStreamingClient(methodInfo serviceinfo.MethodInfo, method string, st streaming.ClientStream) ClientStreamingClient {
	_ = "STUB: not implemented"
	return *new(ClientStreamingClient)
}

func (c *clientStreamingClient) Send(ctx context.Context, req interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *clientStreamingClient) CloseAndRecv(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *clientStreamingClient) Header() (streaming.Header, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Header), nil
}

func (c *clientStreamingClient) Trailer() (streaming.Trailer, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Trailer), nil
}

func (c *clientStreamingClient) CloseSend(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *clientStreamingClient) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *clientStreamingClient) Streaming() streaming.ClientStream {
	_ = "STUB: not implemented"
	return *new(streaming.ClientStream)
}

type ServerStreamingClient interface {
	Recv(ctx context.Context) (interface{}, error)

	Header() (streaming.Header, error)

	Trailer() (streaming.Trailer, error)

	CloseSend(ctx context.Context) error

	Context() context.Context

	Streaming() streaming.ClientStream
}

type serverStreamingClient struct {
	methodInfo serviceinfo.MethodInfo
	method     string
	streaming  streaming.ClientStream
}

func newServerStreamingClient(methodInfo serviceinfo.MethodInfo, method string, st streaming.ClientStream) ServerStreamingClient {
	_ = "STUB: not implemented"
	return *new(ServerStreamingClient)
}

func (c *serverStreamingClient) Recv(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *serverStreamingClient) Header() (streaming.Header, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Header), nil
}

func (c *serverStreamingClient) Trailer() (streaming.Trailer, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Trailer), nil
}

func (c *serverStreamingClient) CloseSend(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *serverStreamingClient) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *serverStreamingClient) Streaming() streaming.ClientStream {
	_ = "STUB: not implemented"
	return *new(streaming.ClientStream)
}

type BidiStreamingClient interface {
	Send(ctx context.Context, req interface{}) error

	Recv(ctx context.Context) (interface{}, error)

	Header() (streaming.Header, error)

	Trailer() (streaming.Trailer, error)

	CloseSend(ctx context.Context) error

	Context() context.Context

	Streaming() streaming.ClientStream
}

type bidiStreamingClient struct {
	methodInfo serviceinfo.MethodInfo
	method     string
	streaming  streaming.ClientStream
}

func newBidiStreamingClient(methodInfo serviceinfo.MethodInfo, method string, st streaming.ClientStream) BidiStreamingClient {
	_ = "STUB: not implemented"
	return *new(BidiStreamingClient)
}

func (c *bidiStreamingClient) Send(ctx context.Context, req interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *bidiStreamingClient) Recv(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *bidiStreamingClient) Header() (streaming.Header, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Header), nil
}

func (c *bidiStreamingClient) Trailer() (streaming.Trailer, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Trailer), nil
}

func (c *bidiStreamingClient) CloseSend(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *bidiStreamingClient) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *bidiStreamingClient) Streaming() streaming.ClientStream {
	_ = "STUB: not implemented"
	return *new(streaming.ClientStream)
}
