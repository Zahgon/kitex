package grpc

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net"
	"sync"
	"sync/atomic"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type bufferPool struct {
	pool sync.Pool
}

func newBufferPool() *bufferPool { _ = "STUB: not implemented"; return nil }

func (p *bufferPool) get() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

func (p *bufferPool) put(b *bytes.Buffer) { _ = "STUB: not implemented"; return }

type recvMsg struct {
	buffer *bytes.Buffer

	err error
}

type recvBuffer struct {
	c       chan recvMsg
	mu      sync.Mutex
	backlog []recvMsg
	err     error
}

func newRecvBuffer() *recvBuffer { _ = "STUB: not implemented"; return nil }

func (b *recvBuffer) put(r recvMsg) { _ = "STUB: not implemented"; return }

func (b *recvBuffer) load() { _ = "STUB: not implemented"; return }

func (b *recvBuffer) get() <-chan recvMsg { _ = "STUB: not implemented"; return nil }

type recvBufferReader struct {
	closeStream func(error)
	ctx         context.Context
	ctxDone     <-chan struct{}
	recv        *recvBuffer
	last        *bytes.Buffer
	err         error
	freeBuffer  func(*bytes.Buffer)
}

func (r *recvBufferReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *recvBufferReader) read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *recvBufferReader) readClient(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *recvBufferReader) readAdditional(m recvMsg, p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type streamState uint32

const (
	streamActive streamState = iota
	streamWriteDone
	streamReadDone
	streamDone
)

type Stream struct {
	id           uint32
	st           ServerTransport
	ct           *http2Client
	ctx          context.Context
	cancel       cancelWithReason
	done         chan struct{}
	ctxDone      <-chan struct{}
	method       string
	recvCompress string
	sendCompress string
	buf          *recvBuffer
	trReader     io.Reader
	fc           *inFlow
	wq           *writeQuota

	requestRead func(int)

	headerChan       chan struct{}
	headerChanClosed uint32

	headerValid bool

	hdrMu sync.Mutex

	header  metadata.MD
	trailer metadata.MD

	noHeaders bool

	headerSent uint32

	state streamState

	status       *status.Status
	reuseStatus  bool
	bizStatusErr kerrors.BizStatusErrorIface

	bytesReceived uint32
	unprocessed   uint32

	contentSubtype string

	closeStreamErr atomic.Value

	sourceService string

	ri rpcinfo.RPCInfo
}

func (s *Stream) isHeaderSent() bool { _ = "STUB: not implemented"; return false }

func (s *Stream) updateHeaderSent() bool { _ = "STUB: not implemented"; return false }

func (s *Stream) swapState(st streamState) streamState {
	_ = "STUB: not implemented"
	return *new(streamState)
}

func (s *Stream) compareAndSwapState(oldState, newState streamState) bool {
	_ = "STUB: not implemented"
	return false
}

func (s *Stream) getState() streamState { _ = "STUB: not implemented"; return *new(streamState) }

func (s *Stream) waitOnHeader() { _ = "STUB: not implemented"; return }

func (s *Stream) RecvCompress() string { _ = "STUB: not implemented"; return "" }

func (s *Stream) SendCompress() string { _ = "STUB: not implemented"; return "" }

func (s *Stream) SetSendCompress(str string) { _ = "STUB: not implemented"; return }

func (s *Stream) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (s *Stream) Header() (metadata.MD, error) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), nil
}

func (s *Stream) tryGetHeader() (metadata.MD, bool) {
	_ = "STUB: not implemented"
	return *new(metadata.MD), false
}

func (s *Stream) getHeaderValid() bool { _ = "STUB: not implemented"; return false }

func (s *Stream) TrailersOnly() bool { _ = "STUB: not implemented"; return false }

func (s *Stream) Trailer() metadata.MD { _ = "STUB: not implemented"; return *new(metadata.MD) }

func (s *Stream) ContentSubtype() string { _ = "STUB: not implemented"; return "" }

func (s *Stream) Context() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

func (s *Stream) Method() string { _ = "STUB: not implemented"; return "" }

func (s *Stream) Status() *status.Status { _ = "STUB: not implemented"; return nil }

func (s *Stream) SetBizStatusErr(bizStatusErr kerrors.BizStatusErrorIface) {
	_ = "STUB: not implemented"
	return
}

func (s *Stream) BizStatusErr() kerrors.BizStatusErrorIface {
	_ = "STUB: not implemented"
	return *new(kerrors.BizStatusErrorIface)
}

func (s *Stream) SetHeader(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) SendHeader(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) SetTrailer(md metadata.MD) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) write(m recvMsg) { _ = "STUB: not implemented"; return }

func (s *Stream) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) getCloseStreamErr() error { _ = "STUB: not implemented"; return nil }

func StreamWrite(s *Stream, buffer *bytes.Buffer) { _ = "STUB: not implemented"; return }

func CreateStream(ctx context.Context, id uint32, requestRead func(i int), method string) *Stream {
	_ = "STUB: not implemented"
	return nil
}

type transportReader struct {
	reader io.Reader

	windowHandler func(int)
	er            error
}

func (t *transportReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Stream) BytesReceived() bool { _ = "STUB: not implemented"; return false }

func (s *Stream) Unprocessed() bool { _ = "STUB: not implemented"; return false }

type transportState int

const (
	reachable transportState = iota
	closing
	draining
)

type ServerConfig struct {
	MaxStreams                 uint32
	KeepaliveParams            ServerKeepalive
	KeepaliveEnforcementPolicy EnforcementPolicy
	InitialWindowSize          uint32
	InitialConnWindowSize      uint32
	WriteBufferSize            uint32
	ReadBufferSize             uint32
	MaxHeaderListSize          *uint32

	ReuseWriteBufferConfig ReuseWriteBufferConfig
}

func DefaultServerConfig() *ServerConfig { _ = "STUB: not implemented"; return nil }

type ConnectOptions struct {
	KeepaliveParams ClientKeepalive

	InitialWindowSize uint32

	InitialConnWindowSize uint32

	WriteBufferSize uint32

	ReadBufferSize uint32

	MaxHeaderListSize *uint32

	ShortConn bool

	TLSConfig *tls.Config

	TraceController *rpcinfo.TraceController

	ReuseWriteBufferConfig ReuseWriteBufferConfig
}

type ClientConfig struct {
	RemoteService string
	OnGoAway      func(ctx context.Context, trans ClientTransport, reason GoAwayReason)
	OnClose       func(ctx context.Context, trans ClientTransport, err error)
}

func NewServerTransport(ctx context.Context, conn net.Conn, cfg *ServerConfig) (ServerTransport, error) {
	_ = "STUB: not implemented"
	return *new(ServerTransport), nil
}

func NewClientTransport(ctx context.Context, conn net.Conn, opts ConnectOptions,
	remoteService string, onGoAway func(GoAwayReason), onClose func(),
) (ClientTransport, error) {
	_ = "STUB: not implemented"
	return *new(ClientTransport), nil
}

var (
	noopOnClose  = func(ctx context.Context, trans ClientTransport, err error) {}
	noopOnGoAway = func(ctx context.Context, trans ClientTransport, reason GoAwayReason) {}
)

func NewClientTransportWithConfig(ctx context.Context, conn net.Conn, opts ConnectOptions,
	cfg ClientConfig,
) (ClientTransport, error) {
	_ = "STUB: not implemented"
	return *new(ClientTransport), nil
}

type Options struct {
	Last bool
}

type CallHdr struct {
	Host string

	Method string

	SendCompress string

	ContentSubtype string

	PreviousAttempts int
}

type IsActive interface {
	IsActive() bool
}

type ClientTransport interface {
	Close(err error) error

	GracefulClose()

	Write(s *Stream, hdr, data []byte, opts *Options) error

	NewStream(ctx context.Context, callHdr *CallHdr) (*Stream, error)

	CloseStream(stream *Stream, err error)

	Error() <-chan struct{}

	GoAway() <-chan struct{}

	GetGoAwayReason() GoAwayReason

	RemoteAddr() net.Addr
	LocalAddr() net.Addr
}

type ServerTransport interface {
	HandleStreams(func(*Stream), func(context.Context, string) context.Context)

	WriteHeader(s *Stream, md metadata.MD) error

	Write(s *Stream, hdr, data []byte, opts *Options) error

	WriteStatus(s *Stream, st *status.Status) error

	Close() error

	RemoteAddr() net.Addr
	LocalAddr() net.Addr

	Drain()
}

func connectionErrorf(temp bool, e error, format string, a ...interface{}) ConnectionError {
	_ = "STUB: not implemented"
	return *new(ConnectionError)
}

func connectionErrorfWithIgnorable(temp bool, e error, format string, a ...interface{}) ConnectionError {
	_ = "STUB: not implemented"
	return *new(ConnectionError)
}

type ConnectionError struct {
	Desc string
	temp bool
	err  error

	isIgnorable bool
}

func (e ConnectionError) Error() string { _ = "STUB: not implemented"; return "" }

func (e ConnectionError) Temporary() bool { _ = "STUB: not implemented"; return false }

func (e ConnectionError) Origin() error { _ = "STUB: not implemented"; return nil }

func (e ConnectionError) Code() int32 { _ = "STUB: not implemented"; return 0 }

func (e ConnectionError) ignorable() bool { _ = "STUB: not implemented"; return false }

func isIgnorable(rawErr error) bool { _ = "STUB: not implemented"; return false }

var (
	ErrConnClosing = connectionErrorfWithIgnorable(true, nil, "transport is closing")

	errStreamDone       = errors.New("the stream is done")
	errStatusStreamDone = status.Err(codes.Internal, errStreamDone.Error())

	errStreamDrain = status.Err(codes.Unavailable, "the connection is draining")

	statusGoAway = status.New(codes.Unavailable, "the stream is rejected because server is draining the connection")
)

type GoAwayReason uint8

const (
	GoAwayInvalid GoAwayReason = 0

	GoAwayNoReason GoAwayReason = 1

	GoAwayTooManyPings GoAwayReason = 2
)

var (
	statusDeadlineExceeded = status.New(codes.DeadlineExceeded, context.DeadlineExceeded.Error())
	errDeadlineExceeded    = statusDeadlineExceeded.Err()
	statusCanceled         = status.New(codes.Canceled, context.Canceled.Error())
	errCanceled            = statusCanceled.Err()
)

func ContextErr(err error) error { _ = "STUB: not implemented"; return nil }

func cascadeContextErr(err error) error { _ = "STUB: not implemented"; return nil }

func standardContextErr(err error) (error, bool) { _ = "STUB: not implemented"; return nil, false }

func defaultContextErr(err error) error { _ = "STUB: not implemented"; return nil }

func tryMarkAsCascadeCancel(err error) error { _ = "STUB: not implemented"; return nil }

func contextStatusAndErr(err error) (*status.Status, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func IsStreamDoneErr(err error) bool { _ = "STUB: not implemented"; return false }

func TLSConfig(tlsConfig *tls.Config) *tls.Config { _ = "STUB: not implemented"; return nil }

const alpnProtoStrH2 = "h2"

func tlsAppendH2ToALPNProtocols(ps []string) []string { _ = "STUB: not implemented"; return nil }

var (
	sendRSTStreamFrameSuffix       = " [send RSTStream Frame]"
	triggeredByRemoteServiceSuffix = " [triggered by remote service]"
	triggeredByHandlerSideSuffix   = " [triggered by handler side]"
)
