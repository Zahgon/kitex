package protobuf

import (
	"context"

	"github.com/cloudwego/kitex/pkg/remote"
)

const (
	metaInfoFixLen = 8
)

func NewProtobufCodec() remote.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(remote.PayloadCodec)
}

func IsProtobufCodec(c remote.PayloadCodec) bool { _ = "STUB: not implemented"; return false }

type protobufCodec struct{}

func (c protobufCodec) Marshal(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c protobufCodec) Unmarshal(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c protobufCodec) Name() string { _ = "STUB: not implemented"; return "" }

type MessageWriterWithContext interface {
	WritePb(ctx context.Context, method string) (interface{}, error)
}

type MessageReaderWithMethodWithContext interface {
	ReadPb(ctx context.Context, method string, in []byte) error
}

type ProtobufMsgCodec interface {
	Marshal(out []byte) ([]byte, error)
	Unmarshal(in []byte) error
}

func getValidData(methodName string, message remote.Message) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
