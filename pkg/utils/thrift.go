package utils

type ThriftMessageCodec struct{}

func NewThriftMessageCodec() *ThriftMessageCodec { _ = "STUB: not implemented"; return nil }

func (t *ThriftMessageCodec) Encode(method string, msgType any, seqID int32, msg any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *ThriftMessageCodec) Decode(b []byte, msg any) (method string, seqID int32, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (t *ThriftMessageCodec) Serialize(msg any) (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *ThriftMessageCodec) Deserialize(msg any, b []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func MarshalError(method string, err error) []byte { _ = "STUB: not implemented"; return nil }

func UnmarshalError(b []byte) error { _ = "STUB: not implemented"; return nil }
