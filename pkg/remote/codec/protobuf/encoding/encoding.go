package encoding

import (
	"io"
)

const Identity = "identity"

type Compressor interface {
	Compress(w io.Writer) (io.WriteCloser, error)

	Decompress(r io.Reader) (io.Reader, error)

	Name() string
}

var registeredCompressor = make(map[string]Compressor)

func RegisterCompressor(c Compressor) { _ = "STUB: not implemented"; return }

func GetCompressor(name string) Compressor { _ = "STUB: not implemented"; return *new(Compressor) }

func FindCompressorName(cname string) string { _ = "STUB: not implemented"; return "" }

func FindCompressor(cname string) (compressor Compressor, err error) {
	_ = "STUB: not implemented"
	return *new(Compressor), nil
}

type Codec interface {
	Marshal(v interface{}) ([]byte, error)

	Unmarshal(data []byte, v interface{}) error

	Name() string
}

var registeredCodecs = make(map[string]Codec)

func RegisterCodec(codec Codec) { _ = "STUB: not implemented"; return }

func GetCodec(contentSubtype string) Codec { _ = "STUB: not implemented"; return *new(Codec) }
