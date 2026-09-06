package remote

import (
	"context"
)

type MetaHandler interface {
	WriteMeta(ctx context.Context, msg Message) (context.Context, error)
	ReadMeta(ctx context.Context, msg Message) (context.Context, error)
}

type StreamingMetaHandler interface {
	OnConnectStream(ctx context.Context) (context.Context, error)

	OnReadStream(ctx context.Context) (context.Context, error)
}
