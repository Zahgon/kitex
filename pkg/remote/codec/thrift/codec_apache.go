package thrift

import (
	"context"
	"sync"

	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/protocol/thrift"
	"github.com/cloudwego/netpoll"

	"github.com/cloudwego/kitex/pkg/remote"
)

func apacheCodecAvailable(data interface{}) bool { _ = "STUB: not implemented"; return false }

func skipThriftStruct(trans bufiox.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func skipThriftStructSlow(trans bufiox.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func apacheMarshal(out bufiox.Writer, ctx context.Context, method string, msgType remote.MessageType, seqID int32, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func apacheUnmarshal(trans bufiox.Reader, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func apacheMarshalData(data interface{}) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type netpollSkipDecoder struct {
	n int
	b []byte
	r netpoll.Reader
}

var skipdecoderPool = sync.Pool{
	New: func() any { return &netpollSkipDecoder{} },
}

func newNetpollSkipDecoder(r netpoll.Reader) *netpollSkipDecoder {
	_ = "STUB: not implemented"
	return nil
}

func (p *netpollSkipDecoder) Reset(r netpoll.Reader) { _ = "STUB: not implemented"; return }

func (p *netpollSkipDecoder) Release() { _ = "STUB: not implemented"; return }

func (p *netpollSkipDecoder) skipn(n int) error { _ = "STUB: not implemented"; return nil }

func (p *netpollSkipDecoder) lastbyte() byte { _ = "STUB: not implemented"; return 0 }

func (p *netpollSkipDecoder) lastbytes(n int) (b []byte) { _ = "STUB: not implemented"; return nil }

func (p *netpollSkipDecoder) skipnSlow(n int) error { _ = "STUB: not implemented"; return nil }

func (p *netpollSkipDecoder) SkipStruct() (int, error) { _ = "STUB: not implemented"; return 0, nil }

var (
	errDepthLimitExceeded = thrift.NewProtocolException(
		thrift.DEPTH_LIMIT, "depth limit exceeded")

	errDataLength = thrift.NewProtocolException(
		thrift.INVALID_DATA, "invalid data length")
)

var typeToSize = [256]int8{
	thrift.BOOL:   1,
	thrift.BYTE:   1,
	thrift.DOUBLE: 8,
	thrift.I16:    2,
	thrift.I32:    4,
	thrift.I64:    8,
}

func (p *netpollSkipDecoder) skipType(t thrift.TType, maxdepth int) error {
	_ = "STUB: not implemented"
	return nil
}
