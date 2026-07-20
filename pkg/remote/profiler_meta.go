package remote

import (
	"context"

	"github.com/cloudwego/kitex/pkg/profiler"
)

type (
	TransInfoTagging func(ctx context.Context, msg Message) (context.Context, []string)

	MessageTagging func(ctx context.Context, msg Message) (context.Context, []string)
)

var _ MetaHandler = (*profilerMetaHandler)(nil)

func NewProfilerMetaHandler(pr profiler.Profiler, tagging MessageTagging) MetaHandler {
	_ = "STUB: not implemented"
	return *new(MetaHandler)
}

type profilerMetaHandler struct {
	profiler profiler.Profiler
	tagging  MessageTagging
}

func (p *profilerMetaHandler) WriteMeta(ctx context.Context, msg Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

func (p *profilerMetaHandler) ReadMeta(ctx context.Context, msg Message) (context.Context, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
