package grpc

import (
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc/grpcframe"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/status"
)

const (
	http2MaxFrameLen = 16384

	http2InitHeaderTableSize = 4096

	baseContentType = "application/grpc"
)

var (
	ClientPreface = []byte(http2.ClientPreface)

	ClientPrefaceLen = len(ClientPreface)
	http2ErrConvTab  = map[http2.ErrCode]codes.Code{
		http2.ErrCodeNo:                 codes.Internal,
		http2.ErrCodeProtocol:           codes.Internal,
		http2.ErrCodeInternal:           codes.Internal,
		http2.ErrCodeFlowControl:        codes.ResourceExhausted,
		http2.ErrCodeSettingsTimeout:    codes.Internal,
		http2.ErrCodeStreamClosed:       codes.Internal,
		http2.ErrCodeFrameSize:          codes.Internal,
		http2.ErrCodeRefusedStream:      codes.Unavailable,
		http2.ErrCodeCancel:             codes.Canceled,
		http2.ErrCodeCompression:        codes.Internal,
		http2.ErrCodeConnect:            codes.Internal,
		http2.ErrCodeEnhanceYourCalm:    codes.ResourceExhausted,
		http2.ErrCodeInadequateSecurity: codes.PermissionDenied,
		http2.ErrCodeHTTP11Required:     codes.Internal,
		gracefulShutdownCode:            codes.Unavailable,
	}
	statusCodeConvTab = map[codes.Code]http2.ErrCode{
		codes.Internal:          http2.ErrCodeInternal,
		codes.Canceled:          http2.ErrCodeCancel,
		codes.Unavailable:       http2.ErrCodeRefusedStream,
		codes.ResourceExhausted: http2.ErrCodeEnhanceYourCalm,
		codes.PermissionDenied:  http2.ErrCodeInadequateSecurity,
	}

	HTTPStatusConvTab = map[int]codes.Code{

		http.StatusBadRequest: codes.Internal,

		http.StatusUnauthorized: codes.Unauthenticated,

		http.StatusForbidden: codes.PermissionDenied,

		http.StatusNotFound: codes.Unimplemented,

		http.StatusTooManyRequests: codes.Unavailable,

		http.StatusBadGateway: codes.Unavailable,

		http.StatusServiceUnavailable: codes.Unavailable,

		http.StatusGatewayTimeout: codes.Unavailable,
	}
)

type parsedHeaderData struct {
	encoding       string
	acceptEncoding string

	statusGen    *status.Status
	bizStatusErr kerrors.BizStatusErrorIface

	rawStatusCode  *int
	rawStatusMsg   string
	bizStatusCode  *int
	bizStatusExtra map[string]string
	httpStatus     *int

	timeoutSet bool
	timeout    time.Duration
	method     string

	mdata          map[string][]string
	statsTags      []byte
	statsTrace     []byte
	contentSubtype string

	isGRPC         bool
	grpcErr        error
	httpErr        error
	contentTypeErr string
}

type decodeState struct {
	serverSide bool

	data parsedHeaderData
}

func isReservedHeader(hdr string) bool { _ = "STUB: not implemented"; return false }

func isWhitelistedHeader(hdr string) bool { _ = "STUB: not implemented"; return false }

func contentSubtype(contentType string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func contentType(contentSubtype string) string { _ = "STUB: not implemented"; return "" }

func (d *decodeState) status() *status.Status { _ = "STUB: not implemented"; return nil }

func (d *decodeState) bizStatusErr() kerrors.BizStatusErrorIface {
	_ = "STUB: not implemented"
	return *new(kerrors.BizStatusErrorIface)
}

func safeCastInt32(n int) int32 { _ = "STUB: not implemented"; return 0 }

const binHdrSuffix = "-bin"

func encodeBinHeader(v []byte) string { _ = "STUB: not implemented"; return "" }

func decodeBinHeader(v string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func encodeMetadataHeader(k, v string) string { _ = "STUB: not implemented"; return "" }

func decodeMetadataHeader(k, v string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *decodeState) decodeHeader(frame *grpcframe.MetaHeadersFrame) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *decodeState) constructHTTPErrMsg() string { _ = "STUB: not implemented"; return "" }

func (d *decodeState) addMetadata(k, v string) { _ = "STUB: not implemented"; return }

func (d *decodeState) processHeaderField(f hpack.HeaderField) { _ = "STUB: not implemented"; return }

type timeoutUnit uint8

const (
	hour        timeoutUnit = 'H'
	minute      timeoutUnit = 'M'
	second      timeoutUnit = 'S'
	millisecond timeoutUnit = 'm'
	microsecond timeoutUnit = 'u'
	nanosecond  timeoutUnit = 'n'
)

func timeoutUnitToDuration(u timeoutUnit) (d time.Duration, ok bool) {
	_ = "STUB: not implemented"
	return *new(time.Duration), false
}

const maxTimeoutValue int64 = 100000000 - 1

func div(d, r time.Duration) int64 { _ = "STUB: not implemented"; return 0 }

func encodeTimeout(t time.Duration) string { _ = "STUB: not implemented"; return "" }

func decodeTimeout(s string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

const (
	spaceByte   = ' '
	tildeByte   = '~'
	percentByte = '%'
)

func encodeGrpcMessage(msg string) string { _ = "STUB: not implemented"; return "" }

func encodeGrpcMessageUnchecked(msg string) string { _ = "STUB: not implemented"; return "" }

func decodeGrpcMessage(msg string) string { _ = "STUB: not implemented"; return "" }

func decodeGrpcMessageUnchecked(msg string) string { _ = "STUB: not implemented"; return "" }
