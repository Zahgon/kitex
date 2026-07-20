package grpc

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

const dataFrameHeaderLen = 5

var (
	ErrInvalidPayload          = errors.New("grpc invalid payload")
	errWrongGRPCImplementation = errors.New("KITEX: grpc client streaming protocol violation: get <nil>, want <EOF>")
)

type marshaler interface {
	MarshalTo(data []byte) (n int, err error)
	Size() int
}

type protobufV2MsgCodec interface {
	XXX_Unmarshal(b []byte) error
	XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
}

type grpcCodec struct {
	ThriftCodec remote.PayloadCodec
}

type CodecOption func(c *grpcCodec)

func WithThriftCodec(t remote.PayloadCodec) CodecOption {
	_ = "STUB: not implemented"
	return *new(CodecOption)
}

func NewGRPCCodec(opts ...CodecOption) remote.Codec {
	_ = "STUB: not implemented"
	return *new(remote.Codec)
}

func mallocWithFirstByteZeroed(size int) []byte { _ = "STUB: not implemented"; return nil }

func (c *grpcCodec) Encode(ctx context.Context, message remote.Message, out remote.ByteBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *grpcCodec) Decode(ctx context.Context, message remote.Message, in remote.ByteBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *grpcCodec) Name() string { _ = "STUB: not implemented"; return "" }

func isNonServerStreaming(mode serviceinfo.StreamingMode) bool {
	_ = "STUB: not implemented"
	return false
}
