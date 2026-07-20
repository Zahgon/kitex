package ttstream

import "github.com/cloudwego/kitex/pkg/rpcinfo"

type ClientHandlerOption func(cp *clientTransHandler)

type ClientProviderOption = ClientHandlerOption

func WithClientMetaFrameHandler(handler MetaFrameHandler) ClientHandlerOption {
	_ = "STUB: not implemented"
	return *new(ClientHandlerOption)
}

func WithClientHeaderFrameHandler(handler HeaderFrameWriteHandler) ClientHandlerOption {
	_ = "STUB: not implemented"
	return *new(ClientHandlerOption)
}

func WithClientLongConnPool(config LongConnConfig) ClientHandlerOption {
	_ = "STUB: not implemented"
	return *new(ClientHandlerOption)
}

func WithClientShortConnPool() ClientHandlerOption {
	_ = "STUB: not implemented"
	return *new(ClientHandlerOption)
}

func WithClientMuxConnPool(config MuxConnConfig) ClientHandlerOption {
	_ = "STUB: not implemented"
	return *new(ClientHandlerOption)
}

func WithClientTraceController(traceCtl *rpcinfo.TraceController) ClientHandlerOption {
	_ = "STUB: not implemented"
	return *new(ClientHandlerOption)
}
