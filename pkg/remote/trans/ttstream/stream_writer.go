package ttstream

var (
	_ streamWriter = (*clientTransport)(nil)
	_ streamWriter = (*serverTransport)(nil)
)

type streamWriter interface {
	WriteFrame(f *Frame) error
	CloseStream(sid int32) error
}
