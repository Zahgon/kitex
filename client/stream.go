package client

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/endpoint/cep"
	"github.com/cloudwego/kitex/pkg/remote/remotecli"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
)

type Streaming interface {
	Stream(ctx context.Context, method string, request, response interface{}) error
	StreamX(ctx context.Context, method string) (streaming.ClientStream, error)
}

func (kc *kClient) Stream(ctx context.Context, method string, request, response interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (kc *kClient) StreamX(ctx context.Context, method string) (streaming.ClientStream, error) {
	_ = "STUB: not implemented"
	return *new(streaming.ClientStream), nil
}

func (kc *kClient) invokeStreamingEndpoint() (endpoint.Endpoint, error) {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint), nil
}

const (
	recvTimeoutErrTpl = "stream Recv timeout, timeout config=%+v"
)

type stream struct {
	streaming.ClientStream
	grpcStream *grpcStream
	ctx        context.Context
	scm        *remotecli.StreamConnManager
	kc         *kClient
	ri         rpcinfo.RPCInfo

	recv      cep.StreamRecvEndpoint
	recvTmCfg streaming.TimeoutConfig
	send      cep.StreamSendEndpoint

	streamingMode serviceinfo.StreamingMode
	finished      uint32
	isGRPC        bool
}

var (
	_ streaming.GRPCStreamGetter = (*stream)(nil)
	_ streaming.WithDoFinish     = (*stream)(nil)
	_ streaming.WithDoFinish     = (*grpcStream)(nil)
)

func newStream(ctx context.Context, s streaming.ClientStream, scm *remotecli.StreamConnManager, kc *kClient, ri rpcinfo.RPCInfo, mode serviceinfo.StreamingMode,
	sendEP cep.StreamSendEndpoint, recvEP cep.StreamRecvEndpoint, grpcSendEP endpoint.SendEndpoint, grpcRecvEP endpoint.RecvEndpoint,
) *stream {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) Header() (hd streaming.Header, err error) {
	_ = "STUB: not implemented"
	return *new(streaming.Header), nil
}

func (s *stream) RecvMsg(ctx context.Context, m interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) handleStreamRecvEvent(err error) { _ = "STUB: not implemented"; return }

func (s *stream) recvWithTimeout(ctx context.Context, m interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) cancel(err error) { _ = "STUB: not implemented"; return }

func (s *stream) SendMsg(ctx context.Context, m interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *stream) handleStreamSendEvent(err error) { _ = "STUB: not implemented"; return }

func (s *stream) DoFinish(err error) { _ = "STUB: not implemented"; return }

func (s *stream) GetGRPCStream() streaming.Stream {
	_ = "STUB: not implemented"
	return *new(streaming.Stream)
}

func newGRPCStream(st streaming.Stream, sendEP endpoint.SendEndpoint, recvEP endpoint.RecvEndpoint,
	recvTmCfg streaming.TimeoutConfig,
) *grpcStream {
	_ = "STUB: not implemented"
	return nil
}

type grpcStream struct {
	streaming.Stream

	st *stream

	sendEndpoint endpoint.SendEndpoint
	recvEndpoint endpoint.RecvEndpoint
	recvTmCfg    streaming.TimeoutConfig
}

func (s *grpcStream) Header() (md metadata.MD, err error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *grpcStream) RecvMsg(m interface{}) (err error) { _ = "STUB: not implemented"; return nil }

func (s *grpcStream) recvWithTimeout(m interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *grpcStream) SendMsg(m interface{}) (err error) { _ = "STUB: not implemented"; return nil }

func (s *grpcStream) DoFinish(err error) { _ = "STUB: not implemented"; return }

func callWithTimeout(tmCfg streaming.TimeoutConfig, call func() error, cancel func(error)) error {
	_ = "STUB: not implemented"
	return nil
}

func isRPCError(err error) bool { _ = "STUB: not implemented"; return false }

var (
	recvEndpoint cep.StreamRecvEndpoint = func(ctx context.Context, stream streaming.ClientStream, m interface{}) error {
		return stream.RecvMsg(ctx, m)
	}
	sendEndpoint cep.StreamSendEndpoint = func(ctx context.Context, stream streaming.ClientStream, m interface{}) error {
		return stream.SendMsg(ctx, m)
	}
)
