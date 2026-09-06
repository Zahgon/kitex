package nphttp2

import (
	"context"
	"net"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

type grpcServerStream struct {
	sx *serverStream
}

type serverStream struct {
	ctx     context.Context
	rpcInfo rpcinfo.RPCInfo
	conn    *serverConn
	handler remote.TransReadWriter

	grpcStream *grpcServerStream
}

var _ streaming.GRPCStreamGetter = (*serverStream)(nil)

func (s *serverStream) GetGRPCStream() streaming.Stream {
	_ = "STUB: not implemented"
	return *new(streaming.Stream)
}

func newServerStream(ctx context.Context, conn *serverConn, handler remote.TransReadWriter) *serverStream {
	_ = "STUB: not implemented"
	return nil
}

func (s *grpcServerStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *grpcServerStream) Trailer() metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}

func (s *grpcServerStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *grpcServerStream) SendHeader(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (s *grpcServerStream) SetHeader(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (s *grpcServerStream) SetTrailer(md metadata.MD) { _ = "STUB: not implemented"; return }

func (s *grpcServerStream) RecvMsg(m interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *grpcServerStream) SendMsg(m interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *grpcServerStream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) SetHeader(hd streaming.Header) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) SendHeader(hd streaming.Header) error { _ = "STUB: not implemented"; return nil }

func (s *serverStream) SetTrailer(tl streaming.Trailer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStream) RecvMsg(ctx context.Context, m interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *serverStream) SendMsg(ctx context.Context, m interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type grpcClientStream struct {
	sx *clientStream
}

type clientStream struct {
	ctx     context.Context
	rpcInfo rpcinfo.RPCInfo
	svcInfo *serviceinfo.ServiceInfo
	conn    *clientConn
	handler remote.TransReadWriter

	grpcStream *grpcClientStream
}

var _ streaming.GRPCStreamGetter = (*clientStream)(nil)

func (s *clientStream) GetGRPCStream() streaming.Stream {
	_ = "STUB: not implemented"
	return *new(streaming.Stream)
}

func NewClientStream(ctx context.Context, svcInfo *serviceinfo.ServiceInfo, conn net.Conn, handler remote.TransReadWriter) streaming.ClientStream {
	_ = "STUB: not implemented"
	return *new(streaming.ClientStream)
}

func (s *grpcClientStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *grpcClientStream) Trailer() metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}

func (s *grpcClientStream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *grpcClientStream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *grpcClientStream) SendHeader(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (s *grpcClientStream) SetHeader(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (s *grpcClientStream) SetTrailer(md metadata.MD) { _ = "STUB: not implemented"; return }

func (s *grpcClientStream) RecvMsg(m interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *grpcClientStream) SendMsg(m interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *clientStream) Header() (streaming.Header, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Header), nil
}

func (s *clientStream) Trailer() (streaming.Trailer, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Trailer), nil
}

func (s *clientStream) RecvMsg(ctx context.Context, m interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *clientStream) SendMsg(ctx context.Context, m interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *clientStream) CloseSend(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *clientStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *clientStream) CancelWithErr(err error) { _ = "STUB: not implemented"; return }

func streamingHeaderToHTTP2MD(header streaming.Header) metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}

func streamingTrailerToHTTP2MD(trailer streaming.Trailer) metadata.MD {
	_ = "STUB: not implemented"
	return *new(metadata.MD)
}

var handleStreamingMetadataMultipleValues func(k string, v []string) string

func HandleStreamingMetadataMultipleValues(hd func(k string, v []string) string) {
	_ = "STUB: not implemented"
	return
}

func http2MDToStreamingHeader(md metadata.MD) (streaming.Header, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Header), nil
}

func http2MDToStreamingTrailer(md metadata.MD) (streaming.Trailer, error) {
	_ = "STUB: not implemented"
	return *new(streaming.Trailer), nil
}
