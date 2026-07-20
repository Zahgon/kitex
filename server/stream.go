package server

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/endpoint/sep"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

func (s *server) wrapStreamMiddleware() endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

func newStream(ctx context.Context, s streaming.ServerStream, sendEP sep.StreamSendEndpoint, recvEP sep.StreamRecvEndpoint,
	traceCtl *rpcinfo.TraceController, grpcSendEP endpoint.SendEndpoint, grpcRecvEP endpoint.RecvEndpoint,
) *stream {
	_ = "STUB: not implemented"
	return nil
}

type stream struct {
	streaming.ServerStream
	grpcStream *grpcStream
	ctx        context.Context
	ri         rpcinfo.RPCInfo
	traceCtl   *rpcinfo.TraceController

	recv sep.StreamRecvEndpoint
	send sep.StreamSendEndpoint
}

var _ streaming.GRPCStreamGetter = (*stream)(nil)

func (s *stream) GetGRPCStream() streaming.Stream {
	_ = "STUB: not implemented"
	return *new(streaming.Stream)
}

func (s *stream) RecvMsg(ctx context.Context, m interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) handleStreamRecvEvent(err error) { _ = "STUB: not implemented"; return }

func (s *stream) SendMsg(ctx context.Context, m interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) handleStreamSendEvent(err error) { _ = "STUB: not implemented"; return }

func newGRPCStream(st streaming.Stream, sendEP endpoint.SendEndpoint, recvEP endpoint.RecvEndpoint) *grpcStream {
	_ = "STUB: not implemented"
	return nil
}

type grpcStream struct {
	streaming.Stream

	st *stream

	sendEndpoint endpoint.SendEndpoint
	recvEndpoint endpoint.RecvEndpoint
}

func (s *grpcStream) RecvMsg(m interface{}) (err error) { _ = "STUB: not implemented"; return nil }

func (s *grpcStream) SendMsg(m interface{}) (err error) { _ = "STUB: not implemented"; return nil }

type contextStream struct {
	streaming.Stream
	ctx context.Context
}

func (cs contextStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

var (
	_ streaming.ServerStream     = (*gRPCCompatibleServerStream)(nil)
	_ streaming.GRPCStreamGetter = (*gRPCCompatibleServerStream)(nil)
)

type gRPCCompatibleServerStream struct {
	streaming.ServerStream
	st streaming.Stream
}

func (gs gRPCCompatibleServerStream) GetGRPCStream() streaming.Stream {
	_ = "STUB: not implemented"
	return *new(streaming.Stream)
}
