package streaming

import (
	"context"
	"io"

	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
)

type Stream interface {
	SetHeader(metadata.MD) error

	SendHeader(metadata.MD) error

	SetTrailer(metadata.MD)

	Header() (metadata.MD, error)

	Trailer() metadata.MD

	Context() context.Context

	RecvMsg(m interface{}) error

	SendMsg(m interface{}) error

	io.Closer
}

type WithDoFinish interface {
	DoFinish(error)
}

type CloseCallbackRegister interface {
	RegisterCloseCallback(cb func(error))
}

type Args struct {
	ServerStream ServerStream
	ClientStream ClientStream

	Stream Stream
}

type Result struct {
	ServerStream ServerStream
	ClientStream ClientStream

	Stream Stream
}

type GRPCStreamGetter interface {
	GetGRPCStream() Stream
}
