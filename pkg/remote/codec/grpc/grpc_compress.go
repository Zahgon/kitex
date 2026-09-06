package grpc

import (
	"context"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/codec/protobuf/encoding"
)

func getSendCompressor(ctx context.Context) (encoding.Compressor, error) {
	_ = "STUB: not implemented"
	return *new(encoding.Compressor), nil
}

func decodeGRPCFrame(ctx context.Context, in remote.ByteBuffer) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func compress(compressor encoding.Compressor, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decompress(compressor encoding.Compressor, data []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
