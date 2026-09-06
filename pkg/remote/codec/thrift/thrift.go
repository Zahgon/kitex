package thrift

import (
	"context"

	"github.com/cloudwego/gopkg/bufiox"

	"github.com/cloudwego/kitex/internal/generic"
	"github.com/cloudwego/kitex/pkg/remote"
)

type CodecType int

const (
	Basic     CodecType = 0b0000
	FastWrite CodecType = 0b0001
	FastRead  CodecType = 0b0010

	FastReadWrite = FastRead | FastWrite

	FrugalWrite CodecType = 0b0100
	FrugalRead  CodecType = 0b1000

	FrugalReadWrite = FrugalWrite | FrugalRead

	EnableSkipDecoder CodecType = 0b10000
)

var (
	errEncodeMismatchMsgType = remote.NewTransErrorWithMsg(remote.InvalidProtocol,
		"encode failed, codec msg type not match with thriftCodec")
	errDecodeMismatchMsgType = remote.NewTransErrorWithMsg(remote.InvalidProtocol,
		"decode failed, codec msg type not match with thriftCodec")
)

func NewThriftCodec() remote.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(remote.PayloadCodec)
}

func IsThriftCodec(c remote.PayloadCodec) bool { _ = "STUB: not implemented"; return false }

func NewThriftCodecWithConfig(c CodecType) remote.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(remote.PayloadCodec)
}

func NewThriftCodecDisableFastMode(disableFastWrite, disableFastRead bool) remote.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(remote.PayloadCodec)
}

type thriftCodec struct {
	CodecType
}

func (c thriftCodec) IsSet(t CodecType) bool { _ = "STUB: not implemented"; return false }

func (c thriftCodec) IsDataLenDeterministic(dataLen int) bool {
	_ = "STUB: not implemented"
	return false
}

func (c thriftCodec) Marshal(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeGenericThrift(out bufiox.Writer, ctx context.Context, method string, msgType remote.MessageType, seqID int32, msg generic.ThriftWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func (c thriftCodec) Unmarshal(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMessageBeforeDecode(ctx context.Context, message remote.Message, seqID int32, methodName string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c thriftCodec) Name() string { _ = "STUB: not implemented"; return "" }
