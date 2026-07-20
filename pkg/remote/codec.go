package remote

import (
	"context"
)

type Codec interface {
	Encode(ctx context.Context, msg Message, out ByteBuffer) error
	Decode(ctx context.Context, msg Message, in ByteBuffer) error
	Name() string
}

type MetaEncoder interface {
	EncodeMetaAndPayload(ctx context.Context, msg Message, out ByteBuffer, me MetaEncoder) error
	EncodePayload(ctx context.Context, msg Message, out ByteBuffer) error
}

type MetaDecoder interface {
	DecodeMeta(ctx context.Context, msg Message, in ByteBuffer) error
	DecodePayload(ctx context.Context, msg Message, in ByteBuffer) error
}
