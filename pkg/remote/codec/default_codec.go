package codec

import (
	"context"

	"github.com/cloudwego/netpoll"

	"github.com/cloudwego/kitex/pkg/remote"
)

const (
	Size32 = 4
	Size16 = 2
)

const (
	ThriftV1Magic = 0x80010000

	ProtobufV1Magic = 0x90010000

	MagicMask = 0xffff0000
)

var (
	ttHeaderCodec   = ttHeader{}
	meshHeaderCodec = meshHeader{}

	_ remote.Codec       = (*defaultCodec)(nil)
	_ remote.MetaDecoder = (*defaultCodec)(nil)
)

func NewDefaultCodec() remote.Codec { _ = "STUB: not implemented"; return *new(remote.Codec) }

func NewDefaultCodecWithSizeLimit(maxSize int) remote.Codec {
	_ = "STUB: not implemented"
	return *new(remote.Codec)
}

func NewDefaultCodecWithConfig(cfg CodecConfig) remote.Codec {
	_ = "STUB: not implemented"
	return *new(remote.Codec)
}

type CodecConfig struct {
	MaxSize int

	CRC32Check bool

	PayloadValidator PayloadValidator
}

type defaultCodec struct {
	CodecConfig
}

func (c *defaultCodec) EncodePayload(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *defaultCodec) EncodeMetaAndPayload(ctx context.Context, message remote.Message, out remote.ByteBuffer, me remote.MetaEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *defaultCodec) Encode(ctx context.Context, message remote.Message, out remote.ByteBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *defaultCodec) DecodeMeta(ctx context.Context, message remote.Message, in remote.ByteBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *defaultCodec) DecodePayload(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *defaultCodec) Decode(ctx context.Context, message remote.Message, in remote.ByteBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *defaultCodec) Name() string { _ = "STUB: not implemented"; return "" }

func (c *defaultCodec) encodeMetaAndPayloadWithPayloadValidator(ctx context.Context, message remote.Message, out remote.ByteBuffer, me remote.MetaEncoder) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *defaultCodec) encodePayload(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func IsTTHeader(flagBuf []byte) bool { _ = "STUB: not implemented"; return false }

func isMeshHeader(flagBuf []byte) bool { _ = "STUB: not implemented"; return false }

func isProtobufKitex(flagBuf []byte) bool { _ = "STUB: not implemented"; return false }

func isThriftBinary(flagBuf []byte) bool { _ = "STUB: not implemented"; return false }

func isThriftFramedBinary(flagBuf []byte) bool { _ = "STUB: not implemented"; return false }

func checkPayload(flagBuf []byte, message remote.Message, in remote.ByteBuffer, isTTHeader bool, maxPayloadSize int) error {
	_ = "STUB: not implemented"
	return nil
}

func checkPayloadSize(payloadLen, maxSize int) error { _ = "STUB: not implemented"; return nil }

func getWrittenBytes(lb *netpoll.LinkBuffer) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
