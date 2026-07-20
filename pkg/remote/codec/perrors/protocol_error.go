package perrors

type ProtocolError interface {
	error
	TypeId() int
}

const (
	UnknownProtocolError = 0
	InvalidData          = 1
	NegativeSize         = 2
	SizeLimit            = 3
	BadVersion           = 4
	NotImplemented       = 5
	DepthLimit           = 6
)

var InvalidDataLength = NewProtocolErrorWithType(InvalidData, "Invalid data length")

type protocolException struct {
	typeID  int
	message string
	rawErr  error
}

func (p *protocolException) TypeId() int { _ = "STUB: not implemented"; return 0 }

func (p *protocolException) String() string { _ = "STUB: not implemented"; return "" }

func (p *protocolException) Error() string { _ = "STUB: not implemented"; return "" }

func (p *protocolException) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (p *protocolException) Is(target error) bool { _ = "STUB: not implemented"; return false }

func NewProtocolError(err error) error { _ = "STUB: not implemented"; return nil }

func NewProtocolErrorWithErrMsg(err error, errMsg string) error {
	_ = "STUB: not implemented"
	return nil
}

func NewProtocolErrorWithMsg(errMsg string) error { _ = "STUB: not implemented"; return nil }

func NewProtocolErrorWithType(errType int, errMsg string) ProtocolError {
	_ = "STUB: not implemented"
	return *new(ProtocolError)
}

func IsProtocolError(err error) bool { _ = "STUB: not implemented"; return false }
