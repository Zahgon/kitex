package generic

import (
	"context"

	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

type ClientStreamingServer interface {
	Recv(ctx context.Context) (req interface{}, err error)

	SendAndClose(ctx context.Context, res interface{}) error

	SetHeader(hd streaming.Header) error

	SendHeader(hd streaming.Header) error

	SetTrailer(hd streaming.Trailer) error

	Streaming() streaming.ServerStream
}

type clientStreamingServer struct {
	methodInfo serviceinfo.MethodInfo
	streaming  streaming.ServerStream
}

func (s *clientStreamingServer) Recv(ctx context.Context) (req interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *clientStreamingServer) SendAndClose(ctx context.Context, res interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *clientStreamingServer) SetHeader(hd streaming.Header) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *clientStreamingServer) SendHeader(hd streaming.Header) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *clientStreamingServer) SetTrailer(hd streaming.Trailer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *clientStreamingServer) Streaming() streaming.ServerStream {
	_ = "STUB: not implemented"
	return *new(streaming.ServerStream)
}

type ServerStreamingServer interface {
	Send(ctx context.Context, res interface{}) error

	SetHeader(hd streaming.Header) error

	SendHeader(hd streaming.Header) error

	SetTrailer(hd streaming.Trailer) error

	Streaming() streaming.ServerStream
}

type serverStreamingServer struct {
	methodInfo serviceinfo.MethodInfo
	streaming  streaming.ServerStream
}

func (s *serverStreamingServer) Send(ctx context.Context, res interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStreamingServer) SetHeader(hd streaming.Header) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStreamingServer) SendHeader(hd streaming.Header) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStreamingServer) SetTrailer(hd streaming.Trailer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStreamingServer) Streaming() streaming.ServerStream {
	_ = "STUB: not implemented"
	return *new(streaming.ServerStream)
}

type BidiStreamingServer interface {
	Recv(ctx context.Context) (req interface{}, err error)

	Send(ctx context.Context, res interface{}) error

	SetHeader(hd streaming.Header) error

	SendHeader(hd streaming.Header) error

	SetTrailer(hd streaming.Trailer) error

	Streaming() streaming.ServerStream
}

type bidiStreamingServer struct {
	methodInfo serviceinfo.MethodInfo
	streaming  streaming.ServerStream
}

func (s *bidiStreamingServer) Recv(ctx context.Context) (req interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *bidiStreamingServer) Send(ctx context.Context, res interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *bidiStreamingServer) SetHeader(hd streaming.Header) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *bidiStreamingServer) SendHeader(hd streaming.Header) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *bidiStreamingServer) SetTrailer(hd streaming.Trailer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *bidiStreamingServer) Streaming() streaming.ServerStream {
	_ = "STUB: not implemented"
	return *new(streaming.ServerStream)
}
