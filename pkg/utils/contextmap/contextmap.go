package contextmap

import (
	"context"
	"sync"
)

type contextMapKey struct{}

func WithContextMap(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func GetContextMap(ctx context.Context) (m *sync.Map, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}
