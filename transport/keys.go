package transport

type Protocol int

const (
	PurePayload Protocol = 0

	TTHeader Protocol = 1 << iota

	Framed

	HTTP

	GRPC

	HESSIAN2

	TTHeaderStreaming

	GRPCStreaming

	TTHeaderFramed = TTHeader | Framed
)

const Unknown = "Unknown"

func (tp Protocol) String() string { _ = "STUB: not implemented"; return "" }
