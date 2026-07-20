package utils

import (
	"strings"
	"sync"
)

func StringDeepCopy(s string) string { _ = "STUB: not implemented"; return "" }

type StringBuilder struct {
	sync.Mutex
	sb strings.Builder
}

func (b *StringBuilder) WithLocked(f func(sb *strings.Builder) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *StringBuilder) RawStringBuilder() *strings.Builder { _ = "STUB: not implemented"; return nil }

func (b *StringBuilder) String() string { _ = "STUB: not implemented"; return "" }

func (b *StringBuilder) Len() int { _ = "STUB: not implemented"; return 0 }

func (b *StringBuilder) Cap() int { _ = "STUB: not implemented"; return 0 }

func (b *StringBuilder) Reset() { _ = "STUB: not implemented"; return }

func (b *StringBuilder) Grow(n int) { _ = "STUB: not implemented"; return }

func (b *StringBuilder) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *StringBuilder) WriteByte(c byte) error { _ = "STUB: not implemented"; return nil }

func (b *StringBuilder) WriteRune(r rune) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *StringBuilder) WriteString(s string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
