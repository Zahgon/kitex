package streaming

import (
	"context"

	"github.com/cloudwego/kitex/pkg/stats"
)

type (
	Header  map[string]string
	Trailer map[string]string
)

type ClientStream interface {
	SendMsg(ctx context.Context, m any) error

	RecvMsg(ctx context.Context, m any) error

	Header() (Header, error)

	Trailer() (Trailer, error)

	CloseSend(ctx context.Context) error

	Context() context.Context
}

type ServerStream interface {
	SendMsg(ctx context.Context, m any) error

	RecvMsg(ctx context.Context, m any) error

	SetHeader(hd Header) error

	SendHeader(hd Header) error

	SetTrailer(hd Trailer) error
}

type ServerStreamingClient[Res any] interface {
	Recv(ctx context.Context) (*Res, error)
	ClientStream
}

func NewServerStreamingClient[Res any](st ClientStream) ServerStreamingClient[Res] {
	_ = "STUB: not implemented"
	return nil
}

type serverStreamingClientImpl[Res any] struct {
	ClientStream
}

func (s *serverStreamingClientImpl[Res]) Recv(ctx context.Context) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ServerStreamingServer[Res any] interface {
	Send(ctx context.Context, res *Res) error
	ServerStream
}

func NewServerStreamingServer[Res any](st ServerStream) ServerStreamingServer[Res] {
	_ = "STUB: not implemented"
	return nil
}

type serverStreamingServerImpl[Res any] struct {
	ServerStream
}

func (s *serverStreamingServerImpl[Res]) Send(ctx context.Context, res *Res) error {
	_ = "STUB: not implemented"
	return nil
}

type ClientStreamingClient[Req, Res any] interface {
	Send(ctx context.Context, req *Req) error
	CloseAndRecv(ctx context.Context) (*Res, error)
	ClientStream
}

func NewClientStreamingClient[Req, Res any](st ClientStream) ClientStreamingClient[Req, Res] {
	_ = "STUB: not implemented"
	return nil
}

type clientStreamingClientImpl[Req, Res any] struct {
	ClientStream
}

func (s *clientStreamingClientImpl[Req, Res]) Send(ctx context.Context, req *Req) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *clientStreamingClientImpl[Req, Res]) CloseAndRecv(ctx context.Context) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ClientStreamingServer[Req, Res any] interface {
	Recv(ctx context.Context) (*Req, error)
	SendAndClose(ctx context.Context, res *Res) error
	ServerStream
}

func NewClientStreamingServer[Req, Res any](st ServerStream) ClientStreamingServer[Req, Res] {
	_ = "STUB: not implemented"
	return nil
}

type clientStreamingServerImpl[Req, Res any] struct {
	ServerStream
}

func (s *clientStreamingServerImpl[Req, Res]) Recv(ctx context.Context) (*Req, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *clientStreamingServerImpl[Req, Res]) SendAndClose(ctx context.Context, res *Res) error {
	_ = "STUB: not implemented"
	return nil
}

type BidiStreamingClient[Req, Res any] interface {
	Send(ctx context.Context, req *Req) error
	Recv(ctx context.Context) (*Res, error)
	ClientStream
}

func NewBidiStreamingClient[Req, Res any](st ClientStream) BidiStreamingClient[Req, Res] {
	_ = "STUB: not implemented"
	return nil
}

type bidiStreamingClientImpl[Req, Res any] struct {
	ClientStream
}

func (s *bidiStreamingClientImpl[Req, Res]) Send(ctx context.Context, req *Req) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *bidiStreamingClientImpl[Req, Res]) Recv(ctx context.Context) (*Res, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type BidiStreamingServer[Req, Res any] interface {
	Recv(ctx context.Context) (*Req, error)
	Send(ctx context.Context, res *Res) error
	ServerStream
}

func NewBidiStreamingServer[Req, Res any](st ServerStream) BidiStreamingServer[Req, Res] {
	_ = "STUB: not implemented"
	return nil
}

type bidiStreamingServerImpl[Req, Res any] struct {
	ServerStream
}

func (s *bidiStreamingServerImpl[Req, Res]) Recv(ctx context.Context) (*Req, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *bidiStreamingServerImpl[Req, Res]) Send(ctx context.Context, res *Res) error {
	_ = "STUB: not implemented"
	return nil
}

type EventHandler func(ctx context.Context, evt stats.Event, err error)
