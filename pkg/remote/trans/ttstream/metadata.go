package ttstream

import (
	"errors"

	"github.com/cloudwego/kitex/pkg/streaming"
)

var (
	ErrInvalidStreamKind = errors.New("invalid stream kind")
	ErrClosedStream      = errors.New("stream is closed")
	ErrCanceledStream    = errors.New("stream is canceled")
)

type (
	IntHeader map[uint16]string
	StrHeader = streaming.Header
)

type ClientStreamMeta interface {
	streaming.ClientStream
	Header() (streaming.Header, error)
	Trailer() (streaming.Trailer, error)
}

type ServerStreamMeta interface {
	streaming.ServerStream
	SetHeader(hd streaming.Header) error
	SendHeader(hd streaming.Header) error
	SetTrailer(hd streaming.Trailer) error
}
