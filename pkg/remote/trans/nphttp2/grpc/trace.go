package grpc

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func (t *http2Client) handleStreamStartEvent(st *Stream, event rpcinfo.StreamStartEvent) {
	_ = "STUB: not implemented"
	return
}

func (t *http2Client) handleStreamRecvHeaderEvent(st *Stream, event rpcinfo.StreamRecvHeaderEvent) {
	_ = "STUB: not implemented"
	return
}

func (t *http2Client) handleStreamFinishEvent(st *Stream, event rpcinfo.StreamFinishEvent) {
	_ = "STUB: not implemented"
	return
}
