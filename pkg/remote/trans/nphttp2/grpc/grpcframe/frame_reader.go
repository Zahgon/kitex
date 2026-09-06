package grpcframe

import (
	"io"

	"github.com/cloudwego/netpoll"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

type Framer struct {
	errDetail error

	lastHeaderStream uint32
	lastFrame        http2.Frame

	reader      netpoll.Reader
	maxReadSize uint32

	writer io.Writer

	reuseWriteBuffer bool
	wbuf             []byte

	AllowIllegalWrites bool

	AllowIllegalReads bool

	ReadMetaHeaders *hpack.Decoder

	MaxHeaderListSize uint32

	frameCache *frameCache
}

func (fr *Framer) maxHeaderListSize() uint32 { _ = "STUB: not implemented"; return 0 }

const (
	minMaxFrameSize = 1 << 14
	maxFrameSize    = 1<<24 - 1
)

func (fr *Framer) SetReuseFrames() { _ = "STUB: not implemented"; return }

type frameCache struct {
	dataFrame DataFrame
}

func (fc *frameCache) getDataFrame() *DataFrame { _ = "STUB: not implemented"; return nil }

func NewFramer(w io.Writer, r netpoll.Reader) *Framer { _ = "STUB: not implemented"; return nil }

func (fr *Framer) SetMaxReadFrameSize(v uint32) { _ = "STUB: not implemented"; return }

func (fr *Framer) SetWriteBufferPoolEnabled(enabled bool) { _ = "STUB: not implemented"; return }

func (fr *Framer) ErrorDetail() error { _ = "STUB: not implemented"; return nil }

func (fr *Framer) ReadFrame() (http2.Frame, error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

func (fr *Framer) connError(code http2.ErrCode, reason string) error {
	_ = "STUB: not implemented"
	return nil
}

func (fr *Framer) checkFrameOrder(f http2.Frame) error { _ = "STUB: not implemented"; return nil }

type headersEnder interface {
	HeadersEnded() bool
}

type headersOrContinuation interface {
	headersEnder
	HeaderBlockFragment() []byte
}

func (fr *Framer) maxHeaderStringLen() int { _ = "STUB: not implemented"; return 0 }

func (fr *Framer) readMetaFrame(hf *HeadersFrame) (*MetaHeadersFrame, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validWireHeaderFieldName(v string) bool { _ = "STUB: not implemented"; return false }

func readByte(p []byte) (remain []byte, b byte, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func readUint32(p []byte) (remain []byte, v uint32, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func (fr *Framer) readAndCheckFrameHeader() (http2.FrameHeader, error) {
	_ = "STUB: not implemented"
	return *new(http2.FrameHeader), nil
}
