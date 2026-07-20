package gzip

import (
	"compress/gzip"
	"io"
	"sync"

	"github.com/cloudwego/kitex/pkg/remote/codec/protobuf/encoding"
)

const Name = "gzip"

func init() {
	c := &compressor{}
	c.poolCompressor.New = func() interface{} {
		return &writer{Writer: gzip.NewWriter(io.Discard), pool: &c.poolCompressor}
	}
	encoding.RegisterCompressor(c)
}

type writer struct {
	*gzip.Writer
	pool *sync.Pool
}

func SetLevel(level int) error { _ = "STUB: not implemented"; return nil }

func (c *compressor) Compress(w io.Writer) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (z *writer) Close() error { _ = "STUB: not implemented"; return nil }

type reader struct {
	*gzip.Reader
	pool *sync.Pool
}

func (c *compressor) Decompress(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

func (z *reader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *compressor) DecompressedSize(buf []byte) int { _ = "STUB: not implemented"; return 0 }

func (c *compressor) Name() string { _ = "STUB: not implemented"; return "" }

type compressor struct {
	poolCompressor   sync.Pool
	poolDecompressor sync.Pool
}
