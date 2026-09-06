package rpcinfo

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/stats"
)

type ClientStreamEventHandler struct {
	HandleStreamStartEvent      func(ctx context.Context, ri RPCInfo, evt StreamStartEvent)
	HandleStreamRecvHeaderEvent func(ctx context.Context, ri RPCInfo, evt StreamRecvHeaderEvent)
	HandleStreamRecvEvent       func(ctx context.Context, ri RPCInfo, evt StreamRecvEvent)
	HandleStreamSendEvent       func(ctx context.Context, ri RPCInfo, evt StreamSendEvent)
	HandleStreamFinishEvent     func(ctx context.Context, ri RPCInfo, evt StreamFinishEvent)
}

type ServerStreamEventHandler struct {
	HandleStreamStartEvent  func(ctx context.Context, ri RPCInfo, evt StreamStartEvent)
	HandleStreamRecvEvent   func(ctx context.Context, ri RPCInfo, evt StreamRecvEvent)
	HandleStreamSendEvent   func(ctx context.Context, ri RPCInfo, evt StreamSendEvent)
	HandleStreamFinishEvent func(ctx context.Context, ri RPCInfo, evt StreamFinishEvent)
}

type StreamStartEvent struct{}

type StreamRecvEvent struct {
	Time time.Time
	Err  error
}

type StreamSendEvent struct {
	Time time.Time
	Err  error
}

type StreamRecvHeaderEvent struct {
	GRPCHeader     map[string][]string
	TTStreamHeader map[string]string
}

type StreamFinishEvent struct {
	GRPCTrailer     map[string][]string
	TTStreamTrailer map[string]string
}

func (c *TraceController) AppendClientStreamEventHandler(hdl ClientStreamEventHandler) {
	_ = "STUB: not implemented"
	return
}

func (c *TraceController) AppendServerStreamEventHandler(hdl ServerStreamEventHandler) {
	_ = "STUB: not implemented"
	return
}

func (c *TraceController) HandleStreamStartEvent(ctx context.Context, ri RPCInfo, evt StreamStartEvent) {
	_ = "STUB: not implemented"
	return
}

func (c *TraceController) HandleStreamRecvEvent(ctx context.Context, ri RPCInfo, evt StreamRecvEvent) {
	_ = "STUB: not implemented"
	return
}

func (c *TraceController) HandleStreamSendEvent(ctx context.Context, ri RPCInfo, evt StreamSendEvent) {
	_ = "STUB: not implemented"
	return
}

func (c *TraceController) HandleStreamRecvHeaderEvent(ctx context.Context, ri RPCInfo, evt StreamRecvHeaderEvent) {
	_ = "STUB: not implemented"
	return
}

func (c *TraceController) HandleStreamFinishEvent(ctx context.Context, ri RPCInfo, evt StreamFinishEvent) {
	_ = "STUB: not implemented"
	return
}

func (c *TraceController) handleStreamRecvEventWrapper(reporter StreamEventReporter) func(ctx context.Context, ri RPCInfo, evt StreamRecvEvent) {
	_ = "STUB: not implemented"
	return nil
}

func (c *TraceController) handleStreamSendEventWrapper(reporter StreamEventReporter) func(ctx context.Context, ri RPCInfo, evt StreamSendEvent) {
	_ = "STUB: not implemented"
	return nil
}

func (c *TraceController) commonStreamIOEventWrapper(ctx context.Context, ri RPCInfo,
	reporter StreamEventReporter, evt stats.Event, err error,
) {
	_ = "STUB: not implemented"
	return
}
