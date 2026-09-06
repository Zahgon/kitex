package apache

import (
	"errors"

	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/protocol/thrift/apache"
)

func init() {

	apache.RegisterCheckTStruct(checkTStruct)
	apache.RegisterThriftRead(callThriftRead)
	apache.RegisterThriftWrite(callThriftWrite)
}

var errNotThriftTStruct = errors.New("not thrift.TStruct")

func checkTStruct(v interface{}) error { _ = "STUB: not implemented"; return nil }

func callThriftRead(r bufiox.Reader, v interface{}) error { _ = "STUB: not implemented"; return nil }

func callThriftWrite(w bufiox.Writer, v interface{}) error { _ = "STUB: not implemented"; return nil }
