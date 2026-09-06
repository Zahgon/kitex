package ttstream

type ServerHandlerOption func(pc *svrTransHandler)

type ServerProviderOption = ServerHandlerOption

func WithServerHeaderFrameHandler(handler HeaderFrameReadHandler) ServerHandlerOption {
	_ = "STUB: not implemented"
	return *new(ServerHandlerOption)
}
