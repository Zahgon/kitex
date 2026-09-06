package apache

import "github.com/apache/thrift/lib/go/thrift"

type TStruct = thrift.TStruct

type TProtocol = thrift.TProtocol

type TTransport = thrift.TTransport

const (
	VERSION_MASK = 0xffff0000
	VERSION_1    = 0x80010000
)

var SkipDefaultDepth = thrift.SkipDefaultDepth

type TException = thrift.TException

var (
	PrependError                  = thrift.PrependError
	NewTProtocolExceptionWithType = thrift.NewTProtocolExceptionWithType
)

const (
	UNKNOWN_PROTOCOL_EXCEPTION = 0
	INVALID_DATA               = 1
	NEGATIVE_SIZE              = 2
	SIZE_LIMIT                 = 3
	BAD_VERSION                = 4
	NOT_IMPLEMENTED            = 5
	DEPTH_LIMIT                = 6
)

type TMessageType = thrift.TMessageType

const (
	INVALID_TMESSAGE_TYPE TMessageType = 0
	CALL                  TMessageType = 1
	REPLY                 TMessageType = 2
	EXCEPTION             TMessageType = 3
	ONEWAY                TMessageType = 4
)

type TType = thrift.TType

const (
	STOP   = 0
	VOID   = 1
	BOOL   = 2
	BYTE   = 3
	I08    = 3
	DOUBLE = 4
	I16    = 6
	I32    = 8
	I64    = 10
	STRING = 11
	UTF7   = 11
	STRUCT = 12
	MAP    = 13
	SET    = 14
	LIST   = 15
	UTF8   = 16
	UTF16  = 17
)
