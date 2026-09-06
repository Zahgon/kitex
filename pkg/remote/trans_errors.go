package remote

const (
	UnknownApplicationException = 0
	UnknownMethod               = 1
	InvalidMessageTypeException = 2
	WrongMethodName             = 3
	BadSequenceID               = 4
	MissingResult               = 5
	InternalError               = 6
	ProtocolError               = 7
	InvalidTransform            = 8
	InvalidProtocol             = 9
	UnsupportedClientType       = 10

	UnknownService = 20
	NoServiceName  = 21
)

var defaultTransErrorMessage = map[int32]string{
	UnknownApplicationException: "unknown application exception",
	UnknownMethod:               "unknown method",
	InvalidMessageTypeException: "invalid message type",
	WrongMethodName:             "wrong method name",
	BadSequenceID:               "bad sequence ID",
	MissingResult:               "missing result",
	InternalError:               "unknown internal error",
	ProtocolError:               "unknown protocol error",
	InvalidTransform:            "Invalid transform",
	InvalidProtocol:             "Invalid protocol",
	UnsupportedClientType:       "Unsupported client type",
	UnknownService:              "unknown service",
}

type TransError struct {
	message string
	typeID  int32
	rawErr  error
}

func (e TransError) Error() string { _ = "STUB: not implemented"; return "" }

func (e TransError) TypeID() int32 { _ = "STUB: not implemented"; return 0 }

func (e TransError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e TransError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (e TransError) AppendMessage(extraMsg string) *TransError {
	_ = "STUB: not implemented"
	return nil
}

func NewTransErrorWithMsg(typeID int32, message string) *TransError {
	_ = "STUB: not implemented"
	return nil
}

func NewTransError(typeID int32, err error) *TransError { _ = "STUB: not implemented"; return nil }

type TypeID interface {
	TypeID() int32
}

type TypeId interface {
	TypeId() int32
}
