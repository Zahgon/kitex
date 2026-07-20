package grpcframe

import (
	"errors"

	"golang.org/x/net/http2"
)

const frameHeaderLen = 9

var padZeros = make([]byte, 255)

var (
	errStreamID    = errors.New("invalid stream ID")
	errDepStreamID = errors.New("invalid dependent stream ID")
)

func (fr *Framer) startWrite(ftype http2.FrameType, flags http2.Flags, streamID uint32, payloadLen int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (fr *Framer) endWrite() (err error) { _ = "STUB: not implemented"; return nil }

func (fr *Framer) writeByte(v byte)     { _ = "STUB: not implemented"; return }
func (fr *Framer) writeBytes(v []byte)  { _ = "STUB: not implemented"; return }
func (fr *Framer) writeUint16(v uint16) { _ = "STUB: not implemented"; return }
func (fr *Framer) writeUint32(v uint32) { _ = "STUB: not implemented"; return }

func (fr *Framer) WriteData(streamID uint32, endStream bool, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (fr *Framer) WriteHeaders(p http2.HeadersFrameParam) error {
	_ = "STUB: not implemented"
	return nil
}

func (fr *Framer) WritePriority(streamID uint32, p http2.PriorityParam) error {
	_ = "STUB: not implemented"
	return nil
}

func (fr *Framer) WriteRSTStream(streamID uint32, code http2.ErrCode) error {
	_ = "STUB: not implemented"
	return nil
}

func (fr *Framer) WriteSettings(settings ...http2.Setting) error {
	_ = "STUB: not implemented"
	return nil
}

func (fr *Framer) WriteSettingsAck() error { _ = "STUB: not implemented"; return nil }

func (fr *Framer) WritePushPromise(p http2.PushPromiseParam) error {
	_ = "STUB: not implemented"
	return nil
}

func (fr *Framer) WritePing(ack bool, data [8]byte) error { _ = "STUB: not implemented"; return nil }

func (fr *Framer) WriteGoAway(maxStreamID uint32, code http2.ErrCode, debugData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (fr *Framer) WriteWindowUpdate(streamID, incr uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (fr *Framer) WriteContinuation(streamID uint32, endHeaders bool, headerBlockFragment []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (fr *Framer) WriteRawFrame(t http2.FrameType, flags http2.Flags, streamID uint32, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func validStreamIDOrZero(streamID uint32) bool { _ = "STUB: not implemented"; return false }

func validStreamID(streamID uint32) bool { _ = "STUB: not implemented"; return false }
