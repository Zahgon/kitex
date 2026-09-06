package thrift

import (
	"sync"
)

type typeCodec struct {
	FastCodec bool
	Frugal    bool
	Apache    bool
}

var type2codec sync.Map

func getTypeCodec(data interface{}) typeCodec { _ = "STUB: not implemented"; return *new(typeCodec) }
