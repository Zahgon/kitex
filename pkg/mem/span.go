package mem

const (
	spanCacheSize = 10
	minSpanObject = 128
	maxSpanObject = (minSpanObject << spanCacheSize) - 1
	minSpanClass  = 8
)

type spanCache struct {
	spans [spanCacheSize]*span
}

func NewSpanCache(spanSize int) *spanCache { _ = "STUB: not implemented"; return nil }

func (c *spanCache) Make(n int) []byte { _ = "STUB: not implemented"; return nil }

func (c *spanCache) Copy(buf []byte) (p []byte) { _ = "STUB: not implemented"; return nil }

func NewSpan(size int) *span { _ = "STUB: not implemented"; return nil }

type span struct {
	lock   uint32
	read   uint32
	size   uint32
	buffer []byte
}

func (b *span) Make(_n int) []byte { _ = "STUB: not implemented"; return nil }

func (b *span) Copy(buf []byte) (p []byte) { _ = "STUB: not implemented"; return nil }

func spanClass(size int) int { _ = "STUB: not implemented"; return 0 }
