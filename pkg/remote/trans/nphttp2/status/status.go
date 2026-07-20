package status

import (
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/codes"
)

type Iface interface {
	GRPCStatus() *Status
}

type Status struct {
	s *spb.Status
}

func New(c codes.Code, msg string) *Status { _ = "STUB: not implemented"; return nil }

func Newf(c codes.Code, format string, a ...interface{}) *Status {
	_ = "STUB: not implemented"
	return nil
}

func ErrorProto(s *spb.Status) error { _ = "STUB: not implemented"; return nil }

func FromProto(s *spb.Status) *Status { _ = "STUB: not implemented"; return nil }

func Err(c codes.Code, msg string) error { _ = "STUB: not implemented"; return nil }

func Errorf(c codes.Code, format string, a ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Status) Code() codes.Code { _ = "STUB: not implemented"; return *new(codes.Code) }

func (s *Status) Message() string { _ = "STUB: not implemented"; return "" }

func (s *Status) AppendMessage(extraMsg string) *Status { _ = "STUB: not implemented"; return nil }

func (s *Status) Proto() *spb.Status { _ = "STUB: not implemented"; return nil }

func (s *Status) Err() error { _ = "STUB: not implemented"; return nil }

func (s *Status) WithDetails(details ...proto.Message) (*Status, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Status) Details() []interface{} { _ = "STUB: not implemented"; return nil }

type Error struct {
	e *spb.Status
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) GRPCStatus() *Status { _ = "STUB: not implemented"; return nil }

func (e *Error) Is(target error) bool { _ = "STUB: not implemented"; return false }

func FromError(err error) (s *Status, ok bool) { _ = "STUB: not implemented"; return nil, false }

func Convert(err error) *Status { _ = "STUB: not implemented"; return nil }

func Code(err error) codes.Code { _ = "STUB: not implemented"; return *new(codes.Code) }

func FromContextError(err error) *Status { _ = "STUB: not implemented"; return nil }
