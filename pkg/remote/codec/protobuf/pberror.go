package protobuf

type PBError interface {
	error
	TypeID() int32
	Marshal(out []byte) ([]byte, error)
	Unmarshal(in []byte) error
}

type pbError struct {
	errProto *ErrorProto
}

func NewPbError(typeID int32, message string) PBError {
	_ = "STUB: not implemented"
	return *new(PBError)
}

func (p pbError) Error() string { _ = "STUB: not implemented"; return "" }

func (p *pbError) IsSetError() bool { _ = "STUB: not implemented"; return false }

func (p *pbError) Marshal(out []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *pbError) Unmarshal(in []byte) error { _ = "STUB: not implemented"; return nil }

func (p *pbError) TypeID() int32 { _ = "STUB: not implemented"; return 0 }
