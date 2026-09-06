package ttstream

import (
	"context"

	"github.com/cloudwego/kitex/pkg/streaming"
)

type HeaderFrameHandler interface {
	HeaderFrameReadHandler
	HeaderFrameWriteHandler
}

type HeaderFrameWriteHandler interface {
	OnWriteStream(ctx context.Context) (ihd IntHeader, shd StrHeader, err error)
}

type HeaderFrameReadHandler interface {
	OnReadStream(ctx context.Context, ihd IntHeader, shd StrHeader) (context.Context, error)
}

type MetaFrameHandler interface {
	OnMetaFrame(ctx context.Context, intHeader IntHeader, header streaming.Header, payload []byte) error
}
