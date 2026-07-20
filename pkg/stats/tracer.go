package stats

import (
	"context"
)

type Tracer interface {
	Start(ctx context.Context) context.Context
	Finish(ctx context.Context)
}
