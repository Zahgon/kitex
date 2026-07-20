package descriptor

import (
	"github.com/cloudwego/gopkg/protocol/thrift"
)

type Type byte

const (
	STOP   Type = 0
	VOID   Type = 1
	BOOL   Type = 2
	BYTE   Type = 3
	I08    Type = 3
	DOUBLE Type = 4
	I16    Type = 6
	I32    Type = 8
	I64    Type = 10
	STRING Type = 11
	UTF7   Type = 11
	STRUCT Type = 12
	MAP    Type = 13
	SET    Type = 14
	LIST   Type = 15
	UTF8   Type = 16
	UTF16  Type = 17

	JSON Type = 19
)

var typeNames = map[Type]string{
	STOP:   "STOP",
	VOID:   "VOID",
	BOOL:   "BOOL",
	BYTE:   "BYTE",
	DOUBLE: "DOUBLE",
	I16:    "I16",
	I32:    "I32",
	I64:    "I64",
	STRING: "STRING",
	STRUCT: "STRUCT",
	MAP:    "MAP",
	SET:    "SET",
	LIST:   "LIST",
	UTF8:   "UTF8",
	UTF16:  "UTF16",
}

func (p Type) String() string { _ = "STUB: not implemented"; return "" }

func (p Type) ToThriftTType() thrift.TType { _ = "STUB: not implemented"; return *new(thrift.TType) }

func FromThriftTType(v interface{}) Type { _ = "STUB: not implemented"; return *new(Type) }

type Void struct{}
