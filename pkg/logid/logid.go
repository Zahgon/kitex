package logid

import (
	"context"
	"sync/atomic"
)

const (
	logIDVersion   = "02"
	logIDMaxLength = 21
	maxRandNum     = 1<<24 - 1<<20
)

var (
	logIDGenerator func(ctx context.Context) string
	logIDExtra     atomic.Value
)

func init() {
	SetLogIDGenerator(DefaultLogIDGenerator)
	SetLogIDExtra("")
}

func SetLogIDGenerator(g func(ctx context.Context) string) { _ = "STUB: not implemented"; return }

func SetLogIDExtra(extra string) { _ = "STUB: not implemented"; return }

func loadLogIDExtra() string { _ = "STUB: not implemented"; return "" }

func DefaultLogIDGenerator(ctx context.Context) string { _ = "STUB: not implemented"; return "" }
