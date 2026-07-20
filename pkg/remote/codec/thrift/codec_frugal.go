package thrift

import (
	"github.com/cloudwego/gopkg/bufiox"

	"github.com/cloudwego/kitex/pkg/remote"
)

func frugalAvailable(data interface{}) bool { _ = "STUB: not implemented"; return false }

func frugalMarshal(out bufiox.Writer, methodName string, msgType remote.MessageType,
	seqID int32, data interface{},
) error {
	_ = "STUB: not implemented"
	return nil
}

func frugalUnmarshal(trans bufiox.Reader, data interface{}, dataLen int) error {
	_ = "STUB: not implemented"
	return nil
}

func frugalMarshalData(data interface{}) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
