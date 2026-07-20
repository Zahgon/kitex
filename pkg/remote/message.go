package remote

import (
	"context"
	"sync"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/transport"
)

var (
	messagePool      sync.Pool
	transInfoPool    sync.Pool
	emptyServiceInfo serviceinfo.ServiceInfo
)

func init() {
	messagePool.New = newMessage
	transInfoPool.New = newTransInfo
}

type MessageType int32

const (
	InvalidMessageType MessageType = 0
	Call               MessageType = 1
	Reply              MessageType = 2
	Exception          MessageType = 3

	Oneway MessageType = 4

	Stream MessageType = 5

	Heartbeat MessageType = 6
)

const (
	ReadFailed string = "RFailed"

	MeshHeader string = "mHeader"
)

type ProtocolInfo struct {
	TransProto transport.Protocol
	CodecType  serviceinfo.PayloadCodec
}

type ServiceSearcher interface {
	SearchService(svcName, methodName string, strict bool, codecType serviceinfo.PayloadCodec) *serviceinfo.ServiceInfo
}

type keyServiceSearcher struct{}

func GetServiceSearcher(ctx context.Context) ServiceSearcher {
	_ = "STUB: not implemented"
	return *new(ServiceSearcher)
}

func WithServiceSearcher(ctx context.Context, svcSearcher ServiceSearcher) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type Message interface {
	RPCInfo() rpcinfo.RPCInfo
	Data() interface{}
	NewData(method string) (ok bool)
	MessageType() MessageType
	SetMessageType(MessageType)
	RPCRole() RPCRole
	PayloadLen() int
	SetPayloadLen(size int)
	TransInfo() TransInfo
	Tags() map[string]interface{}
	PayloadCodec() PayloadCodec
	SetPayloadCodec(pc PayloadCodec)
	Recycle()

	ProtocolInfo() ProtocolInfo
}

func NewMessage(data interface{}, ri rpcinfo.RPCInfo, msgType MessageType, rpcRole RPCRole) Message {
	_ = "STUB: not implemented"
	return *new(Message)
}

func RecycleMessage(msg Message) { _ = "STUB: not implemented"; return }

func newMessage() interface{} { _ = "STUB: not implemented"; return nil }

type message struct {
	msgType      MessageType
	data         interface{}
	rpcInfo      rpcinfo.RPCInfo
	rpcRole      RPCRole
	compressType CompressType
	payloadSize  int
	transInfo    TransInfo
	tags         map[string]interface{}
	payloadCodec PayloadCodec
}

func (m *message) zero() { _ = "STUB: not implemented"; return }

func (m *message) RPCInfo() rpcinfo.RPCInfo {
	_ = "STUB: not implemented"
	return *new(rpcinfo.RPCInfo)
}

func (m *message) Data() interface{} { _ = "STUB: not implemented"; return nil }

func (m *message) NewData(method string) (ok bool) { _ = "STUB: not implemented"; return false }

func (m *message) MessageType() MessageType { _ = "STUB: not implemented"; return *new(MessageType) }

func (m *message) SetMessageType(mt MessageType) { _ = "STUB: not implemented"; return }

func (m *message) RPCRole() RPCRole { _ = "STUB: not implemented"; return *new(RPCRole) }

func (m *message) TransInfo() TransInfo { _ = "STUB: not implemented"; return *new(TransInfo) }

func (m *message) Tags() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func (m *message) ProtocolInfo() ProtocolInfo { _ = "STUB: not implemented"; return *new(ProtocolInfo) }

func (m *message) PayloadLen() int { _ = "STUB: not implemented"; return 0 }

func (m *message) SetPayloadLen(size int) { _ = "STUB: not implemented"; return }

func (m *message) PayloadCodec() PayloadCodec { _ = "STUB: not implemented"; return *new(PayloadCodec) }

func (m *message) SetPayloadCodec(pc PayloadCodec) { _ = "STUB: not implemented"; return }

func (m *message) Recycle() { _ = "STUB: not implemented"; return }

type TransInfo interface {
	TransStrInfo() map[string]string
	TransIntInfo() map[uint16]string
	PutTransIntInfo(map[uint16]string)
	PutTransStrInfo(kvInfo map[string]string)
	Recycle()
}

func newTransInfo() interface{} { _ = "STUB: not implemented"; return nil }

type transInfo struct {
	strInfo map[string]string
	intInfo map[uint16]string
}

func (ti *transInfo) zero() { _ = "STUB: not implemented"; return }

func (ti *transInfo) TransIntInfo() map[uint16]string { _ = "STUB: not implemented"; return nil }

func (ti *transInfo) PutTransIntInfo(kvInfo map[uint16]string) { _ = "STUB: not implemented"; return }

func (ti *transInfo) TransStrInfo() map[string]string { _ = "STUB: not implemented"; return nil }

func (ti *transInfo) PutTransStrInfo(kvInfo map[string]string) { _ = "STUB: not implemented"; return }

func (ti *transInfo) Recycle() { _ = "STUB: not implemented"; return }
