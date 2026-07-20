package codec

import (
	"context"

	"github.com/cloudwego/kitex/pkg/remote"
)

const (
	FrontMask = 0x0000ffff
)

func SetOrCheckMethodName(ctx context.Context, methodName string, message remote.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func SetOrCheckSeqID(seqID int32, message remote.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func UpdateMsgType(msgType uint32, message remote.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataIfNeeded(method string, message remote.Message) error {
	_ = "STUB: not implemented"
	return nil
}
