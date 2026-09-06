package rpcinfo

import (
	"context"

	"github.com/cloudwego/kitex/internal/stream"
	"github.com/cloudwego/kitex/pkg/stats"
)

type StreamEventReporter interface {
	ReportStreamEvent(ctx context.Context, ri RPCInfo, event Event)
}

type TraceController struct {
	tracers              []stats.Tracer
	streamEventReporters []StreamEventReporter

	streamStartEventHandlers      []func(ctx context.Context, ri RPCInfo, evt StreamStartEvent)
	streamRecvEventHandlers       []func(ctx context.Context, ri RPCInfo, evt StreamRecvEvent)
	streamSendEventHandlers       []func(ctx context.Context, ri RPCInfo, evt StreamSendEvent)
	streamRecvHeaderEventHandlers []func(ctx context.Context, ri RPCInfo, evt StreamRecvHeaderEvent)
	streamFinishEventHandlers     []func(ctx context.Context, ri RPCInfo, evt StreamFinishEvent)
}

func (c *TraceController) Append(col stats.Tracer) { _ = "STUB: not implemented"; return }

func (c *TraceController) DoStart(ctx context.Context, ri RPCInfo) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (c *TraceController) DoFinish(ctx context.Context, ri RPCInfo, err error) {
	_ = "STUB: not implemented"
	return
}

func buildStreamingEvent(statsEvent stats.Event, err error) Event {
	_ = "STUB: not implemented"
	return *new(Event)
}

func (c *TraceController) ReportStreamEvent(ctx context.Context, statsEvent stats.Event, err error) {
	_ = "STUB: not implemented"
	return
}

func (c *TraceController) GetStreamEventHandler() stream.StreamEventHandler {
	_ = "STUB: not implemented"
	return *new(stream.StreamEventHandler)
}

func (c *TraceController) HasTracer() bool { _ = "STUB: not implemented"; return false }

func (c *TraceController) HasStreamEventReporter() bool { _ = "STUB: not implemented"; return false }

func (c *TraceController) tryRecover(ctx context.Context) { _ = "STUB: not implemented"; return }
