package generic

import (
	"context"

	"github.com/cloudwego/kitex/pkg/generic/thrift"
	"github.com/cloudwego/kitex/pkg/remote"
)

var (
	_  remote.PayloadCodec = &binaryThriftCodec{}
	wb                     = thrift.NewWriteBinary()
)

type binaryReqType = []byte

type binaryThriftCodec struct {
	thriftCodec remote.PayloadCodec
}

func (c *binaryThriftCodec) Marshal(ctx context.Context, msg remote.Message, out remote.ByteBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *binaryThriftCodec) Unmarshal(ctx context.Context, msg remote.Message, in remote.ByteBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *binaryThriftCodec) Name() string { _ = "STUB: not implemented"; return "" }

func SetSeqID(seqID int32, transBuff []byte) error { _ = "STUB: not implemented"; return nil }

func GetSeqID(transBuff []byte) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

func getSeqID4Bytes(transBuff []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func readBinaryMethod(ctx context.Context, buff []byte, msg remote.Message) error {
	_ = "STUB: not implemented"
	return nil
}
