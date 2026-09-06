package codec

import (
	"github.com/cloudwego/kitex/pkg/remote"
)

func ReadUint32(in remote.ByteBuffer) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func PeekUint32(in remote.ByteBuffer) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func ReadUint16(in remote.ByteBuffer) (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func WriteUint32(val uint32, out remote.ByteBuffer) error { _ = "STUB: not implemented"; return nil }

func WriteUint16(val uint16, out remote.ByteBuffer) error { _ = "STUB: not implemented"; return nil }

func WriteByte(val byte, out remote.ByteBuffer) error { _ = "STUB: not implemented"; return nil }

func Bytes2Uint32NoCheck(bytes []byte) uint32 { _ = "STUB: not implemented"; return 0 }

func Bytes2Uint32(bytes []byte) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func Bytes2Uint16NoCheck(bytes []byte) uint16 { _ = "STUB: not implemented"; return 0 }

func Bytes2Uint16(bytes []byte, off int) (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func Bytes2Uint8(bytes []byte, off int) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

func ReadString(in remote.ByteBuffer) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func WriteString(val string, out remote.ByteBuffer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func WriteString2BLen(val string, out remote.ByteBuffer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func ReadString2BLen(bytes []byte, off int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}
