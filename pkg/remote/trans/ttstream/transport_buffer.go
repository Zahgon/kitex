package ttstream

import (
	"sync"

	"github.com/cloudwego/gopkg/bufiox"
	gopkgthrift "github.com/cloudwego/gopkg/protocol/thrift"
	"github.com/cloudwego/netpoll"
)

var (
	_ bufiox.Reader            = (*readerBuffer)(nil)
	_ bufiox.Writer            = (*writerBuffer)(nil)
	_ gopkgthrift.NocopyWriter = (*writerBuffer)(nil)
)

var (
	readerBufferPool sync.Pool
	writerBufferPool sync.Pool
)

func newReaderBuffer(reader netpoll.Reader) (rb *readerBuffer) {
	_ = "STUB: not implemented"
	return nil
}

type readerBuffer struct {
	reader   netpoll.Reader
	readSize int
}

func (c *readerBuffer) Next(n int) (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *readerBuffer) ReadBinary(bs []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *readerBuffer) Peek(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *readerBuffer) Skip(n int) (err error) { _ = "STUB: not implemented"; return nil }

func (c *readerBuffer) ReadLen() (n int) { _ = "STUB: not implemented"; return 0 }

func (c *readerBuffer) Release(e error) (err error) { _ = "STUB: not implemented"; return nil }

func newWriterBuffer(writer netpoll.Writer) (wb *writerBuffer) {
	_ = "STUB: not implemented"
	return nil
}

type writerBuffer struct {
	writer    netpoll.Writer
	writeSize int
}

func (c *writerBuffer) Malloc(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *writerBuffer) WriteBinary(bs []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *writerBuffer) WriteDirect(b []byte, remainCap int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *writerBuffer) WrittenLen() (length int) { _ = "STUB: not implemented"; return 0 }

func (c *writerBuffer) Flush() (err error) { _ = "STUB: not implemented"; return nil }
