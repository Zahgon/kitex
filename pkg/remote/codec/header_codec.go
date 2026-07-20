package codec

import (
	"context"

	"github.com/cloudwego/gopkg/protocol/ttheader"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

const (
	TTHeaderMagic            = ttheader.TTHeaderMagic
	MeshHeaderMagic   uint32 = 0xFFAF0000
	MeshHeaderLenMask uint32 = 0x0000FFFF

	FlagsMask     = ttheader.FlagsMask
	MethodMask    = ttheader.MethodMask
	MaxFrameSize  = ttheader.MaxFrameSize
	MaxHeaderSize = ttheader.MaxHeaderSize
)

type HeaderFlags = ttheader.HeaderFlags

const (
	HeaderFlagsKey              string = "HeaderFlags"
	HeaderFlagSupportOutOfOrder        = ttheader.HeaderFlagSupportOutOfOrder
	HeaderFlagDuplexReverse            = ttheader.HeaderFlagDuplexReverse
	HeaderFlagSASL                     = ttheader.HeaderFlagSASL
)

const (
	TTHeaderMetaSize = ttheader.TTHeaderMetaSize
)

type ProtocolID = ttheader.ProtocolID

const (
	ProtocolIDThriftBinary    = ttheader.ProtocolIDThriftBinary
	ProtocolIDThriftCompact   = ttheader.ProtocolIDThriftCompact
	ProtocolIDThriftCompactV2 = ttheader.ProtocolIDThriftCompactV2
	ProtocolIDKitexProtobuf   = ttheader.ProtocolIDKitexProtobuf
	ProtocolIDDefault         = ttheader.ProtocolIDDefault
)

type InfoIDType = ttheader.InfoIDType

const (
	InfoIDPadding     = ttheader.InfoIDPadding
	InfoIDKeyValue    = ttheader.InfoIDKeyValue
	InfoIDIntKeyValue = ttheader.InfoIDIntKeyValue
	InfoIDACLToken    = ttheader.InfoIDACLToken
)

type ttHeader struct{}

func (t ttHeader) encode(ctx context.Context, message remote.Message, out remote.ByteBuffer) (totalLenField []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t ttHeader) decode(ctx context.Context, message remote.Message, in remote.ByteBuffer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func readStrKVInfo(idx *int, buf []byte, info map[string]string) (has bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getFlags(message remote.Message) ttheader.HeaderFlags {
	_ = "STUB: not implemented"
	return *new(ttheader.HeaderFlags)
}

func setFlags(flags ttheader.HeaderFlags, message remote.Message) {
	_ = "STUB: not implemented"
	return
}

func getProtocolID(ct serviceinfo.PayloadCodec) ProtocolID {
	_ = "STUB: not implemented"
	return *new(ProtocolID)
}

type meshHeader struct{}

//lint:ignore U1000 until encode is used
func (m meshHeader) encode(ctx context.Context, message remote.Message, payloadBuf, out remote.ByteBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (m meshHeader) decode(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func fillBasicInfoOfTTHeader(msg remote.Message) { _ = "STUB: not implemented"; return }
