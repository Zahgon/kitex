package metadata

import (
	"context"
)

func DecodeKeyValue(k, v string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type MD map[string][]string

func New(m map[string]string) MD { _ = "STUB: not implemented"; return *new(MD) }

func Pairs(kv ...string) MD { _ = "STUB: not implemented"; return *new(MD) }

func (md MD) Len() int { _ = "STUB: not implemented"; return 0 }

func (md MD) Copy() MD { _ = "STUB: not implemented"; return *new(MD) }

func (md MD) Get(k string) []string { _ = "STUB: not implemented"; return nil }

func (md MD) Set(k string, vals ...string) { _ = "STUB: not implemented"; return }

func (md MD) Append(k string, vals ...string) { _ = "STUB: not implemented"; return }

func Join(mds ...MD) MD { _ = "STUB: not implemented"; return *new(MD) }

func AppendMD(md, other MD) MD { _ = "STUB: not implemented"; return *new(MD) }

type (
	mdIncomingKey struct{}
	mdOutgoingKey struct{}
)

func NewIncomingContext(ctx context.Context, md MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func NewOutgoingContext(ctx context.Context, md MD) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func AppendToOutgoingContext(ctx context.Context, kv ...string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromIncomingContext(ctx context.Context) (md MD, ok bool) {
	_ = "STUB: not implemented"
	return *new(MD), false
}

func FromOutgoingContextRaw(ctx context.Context) (MD, [][]string, bool) {
	_ = "STUB: not implemented"
	return *new(MD), nil, false
}

func FromOutgoingContext(ctx context.Context) (MD, bool) {
	_ = "STUB: not implemented"
	return *new(MD), false
}

type rawMD struct {
	md    MD
	added [][]string
}
