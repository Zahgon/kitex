package remote

import (
	"io"

	"github.com/cloudwego/gopkg/bufiox"
)

type bufioxBuffer struct {
	io.ReadWriter
	bufiox.Writer
	bufiox.Reader
}

func NewByteBufferFromBufiox(bw bufiox.Writer, br bufiox.Reader) ByteBuffer {
	_ = "STUB: not implemented"
	return *new(ByteBuffer)
}

func (b *bufioxBuffer) ReadableLen() (n int) { _ = "STUB: not implemented"; return 0 }

func (b *bufioxBuffer) ReadString(n int) (s string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *bufioxBuffer) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *bufioxBuffer) NewBuffer() ByteBuffer { _ = "STUB: not implemented"; return *new(ByteBuffer) }

func (b *bufioxBuffer) AppendBuffer(buf ByteBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *bufioxBuffer) Bytes() (buf []byte, err error) { _ = "STUB: not implemented"; return nil, nil }
