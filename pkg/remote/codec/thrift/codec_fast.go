package thrift

import (
	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/protocol/thrift"

	"github.com/cloudwego/kitex/pkg/remote"
)

type ThriftMsgFastCodec = thrift.FastCodec

func fastCodecAvailable(data interface{}) bool { _ = "STUB: not implemented"; return false }

func fastMarshal(out bufiox.Writer, methodName string, msgType remote.MessageType, seqID int32, msg thrift.FastCodec) error {
	_ = "STUB: not implemented"
	return nil
}

func fastUnmarshal(trans bufiox.Reader, data interface{}, dataLen int) error {
	_ = "STUB: not implemented"
	return nil
}

func fastMarshalData(data interface{}) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
