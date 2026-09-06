package remote

import (
	"context"

	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

var payloadCodecs = make(map[serviceinfo.PayloadCodec]PayloadCodec)

type PayloadCodec interface {
	Marshal(ctx context.Context, message Message, out ByteBuffer) error
	Unmarshal(ctx context.Context, message Message, in ByteBuffer) error
	Name() string
}

func GetPayloadCodec(message Message) (PayloadCodec, error) {
	_ = "STUB: not implemented"
	return *new(PayloadCodec), nil
}

func PutPayloadCode(name serviceinfo.PayloadCodec, v PayloadCodec) {
	_ = "STUB: not implemented"
	return
}
