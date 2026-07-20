package grpcframe

import (
	"github.com/cloudwego/netpoll"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

type frameParser func(fc *frameCache, fh http2.FrameHeader, payload []byte) (http2.Frame, error)

var frameParsers = []frameParser{

	http2.FrameHeaders:      parseHeadersFrame,
	http2.FramePriority:     parsePriorityFrame,
	http2.FrameRSTStream:    parseRSTStreamFrame,
	http2.FrameSettings:     parseSettingsFrame,
	http2.FramePushPromise:  parsePushPromise,
	http2.FramePing:         parsePingFrame,
	http2.FrameGoAway:       parseGoAwayFrame,
	http2.FrameWindowUpdate: parseWindowUpdateFrame,
	http2.FrameContinuation: parseContinuationFrame,
}

func typeFrameParser(t http2.FrameType) frameParser {
	_ = "STUB: not implemented"
	return *new(frameParser)
}

type DataFrame struct {
	http2.FrameHeader
	data []byte
}

func (f *DataFrame) StreamEnded() bool { _ = "STUB: not implemented"; return false }

func (f *DataFrame) Data() []byte { _ = "STUB: not implemented"; return nil }

func parseDataFrame(fc *frameCache, fh http2.FrameHeader, payload netpoll.Reader) (http2.Frame, error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

type HeadersFrame struct {
	http2.FrameHeader

	Priority http2.PriorityParam

	headerFragBuf []byte
}

func (f *HeadersFrame) HeaderBlockFragment() []byte { _ = "STUB: not implemented"; return nil }

func (f *HeadersFrame) HeadersEnded() bool { _ = "STUB: not implemented"; return false }

func (f *HeadersFrame) StreamEnded() bool { _ = "STUB: not implemented"; return false }

func (f *HeadersFrame) HasPriority() bool { _ = "STUB: not implemented"; return false }

func parseHeadersFrame(_ *frameCache, fh http2.FrameHeader, p []byte) (_ http2.Frame, err error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

func parsePriorityFrame(_ *frameCache, fh http2.FrameHeader, payload []byte) (http2.Frame, error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

func parseRSTStreamFrame(_ *frameCache, fh http2.FrameHeader, p []byte) (http2.Frame, error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

type SettingsFrame struct {
	http2.FrameHeader
	p []byte
}

func parseSettingsFrame(_ *frameCache, fh http2.FrameHeader, p []byte) (http2.Frame, error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

func (f *SettingsFrame) IsAck() bool { _ = "STUB: not implemented"; return false }

func (f *SettingsFrame) Value(id http2.SettingID) (v uint32, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func (f *SettingsFrame) Setting(i int) http2.Setting {
	_ = "STUB: not implemented"
	return *new(http2.Setting)
}

func (f *SettingsFrame) NumSettings() int { _ = "STUB: not implemented"; return 0 }

func (f *SettingsFrame) HasDuplicates() bool { _ = "STUB: not implemented"; return false }

func (f *SettingsFrame) ForeachSetting(fn func(http2.Setting) error) error {
	_ = "STUB: not implemented"
	return nil
}

type PushPromiseFrame struct {
	http2.FrameHeader
	PromiseID     uint32
	headerFragBuf []byte
}

func (f *PushPromiseFrame) HeaderBlockFragment() []byte { _ = "STUB: not implemented"; return nil }

func (f *PushPromiseFrame) HeadersEnded() bool { _ = "STUB: not implemented"; return false }

func parsePushPromise(_ *frameCache, fh http2.FrameHeader, p []byte) (_ http2.Frame, err error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

func parsePingFrame(_ *frameCache, fh http2.FrameHeader, payload []byte) (http2.Frame, error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

type GoAwayFrame struct {
	http2.FrameHeader
	LastStreamID uint32
	ErrCode      http2.ErrCode
	debugData    []byte
}

func (f *GoAwayFrame) DebugData() []byte { _ = "STUB: not implemented"; return nil }

func parseGoAwayFrame(_ *frameCache, fh http2.FrameHeader, p []byte) (http2.Frame, error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

func parseWindowUpdateFrame(_ *frameCache, fh http2.FrameHeader, p []byte) (http2.Frame, error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

type ContinuationFrame struct {
	http2.FrameHeader
	headerFragBuf []byte
}

func parseContinuationFrame(_ *frameCache, fh http2.FrameHeader, p []byte) (http2.Frame, error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

func (f *ContinuationFrame) HeaderBlockFragment() []byte { _ = "STUB: not implemented"; return nil }

func (f *ContinuationFrame) HeadersEnded() bool { _ = "STUB: not implemented"; return false }

type UnknownFrame struct {
	http2.FrameHeader
	p []byte
}

func (f *UnknownFrame) Payload() []byte { _ = "STUB: not implemented"; return nil }

func parseUnknownFrame(_ *frameCache, fh http2.FrameHeader, p []byte) (http2.Frame, error) {
	_ = "STUB: not implemented"
	return *new(http2.Frame), nil
}

type MetaHeadersFrame struct {
	*HeadersFrame

	Fields []hpack.HeaderField

	Truncated bool
}

func (mh *MetaHeadersFrame) PseudoValue(pseudo string) string { _ = "STUB: not implemented"; return "" }

func (mh *MetaHeadersFrame) RegularFields() []hpack.HeaderField {
	_ = "STUB: not implemented"
	return nil
}

func (mh *MetaHeadersFrame) PseudoFields() []hpack.HeaderField {
	_ = "STUB: not implemented"
	return nil
}

func (mh *MetaHeadersFrame) checkPseudos() error { _ = "STUB: not implemented"; return nil }
