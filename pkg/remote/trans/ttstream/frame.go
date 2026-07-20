package ttstream

import (
	"context"
	"errors"
	"sync"

	"github.com/cloudwego/gopkg/bufiox"
	gopkgthrift "github.com/cloudwego/gopkg/protocol/thrift"
	"github.com/cloudwego/gopkg/protocol/ttheader"
)

const (
	metaFrameType    int32 = 1
	headerFrameType  int32 = 2
	dataFrameType    int32 = 3
	trailerFrameType int32 = 4
	rstFrameType     int32 = 5
)

var frameTypeToString = map[int32]string{
	metaFrameType:    ttheader.FrameTypeMeta,
	headerFrameType:  ttheader.FrameTypeHeader,
	dataFrameType:    ttheader.FrameTypeData,
	trailerFrameType: ttheader.FrameTypeTrailer,
	rstFrameType:     ttheader.FrameTypeRst,
}

var framePool sync.Pool

var (
	errNoRPCInfo      = errors.New("no rpcinfo in context")
	errInvalidMessage = errors.New("ttheaderstreaming invalid message")
)

type Frame struct {
	streamFrame
	typ     int32
	payload []byte
}

func (f *Frame) String() string { _ = "STUB: not implemented"; return "" }

func newFrame(sframe streamFrame, typ int32, payload []byte) (fr *Frame) {
	_ = "STUB: not implemented"
	return nil
}

func recycleFrame(frame *Frame) { _ = "STUB: not implemented"; return }

func EncodeFrame(ctx context.Context, writer bufiox.Writer, fr *Frame) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func DecodeFrame(ctx context.Context, reader bufiox.Reader) (fr *Frame, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodePayload(ctx context.Context, msg any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DecodePayload(ctx context.Context, payload []byte, msg any) error {
	_ = "STUB: not implemented"
	return nil
}

func EncodeException(ctx context.Context, method string, seq int32, ex error) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getThriftMessageTypeStr(typ gopkgthrift.TMessageType) string {
	_ = "STUB: not implemented"
	return ""
}

func decodeException(buf []byte) (*gopkgthrift.ApplicationException, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeFrameAndFlush(ctx context.Context, writer bufiox.Writer, fr *Frame) (err error) {
	_ = "STUB: not implemented"
	return nil
}
