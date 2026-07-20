package generic

import (
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/codec/protobuf"
	"github.com/cloudwego/kitex/pkg/remote/codec/thrift"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

type Generic interface {
	Closer

	PayloadCodecType() serviceinfo.PayloadCodec

	GenericMethod() serviceinfo.GenericMethodFunc

	IDLServiceName() string

	GetExtra(key string) interface{}
}

type GetMethodNameByRequestFunc func(req interface{}) (string, error)

type messageReaderWriterGetter interface {
	getMessageReaderWriter() interface{}
}

type Method struct {
	Oneway        bool
	StreamingMode serviceinfo.StreamingMode
}

func BinaryThriftGeneric() Generic { _ = "STUB: not implemented"; return *new(Generic) }

func BinaryThriftGenericV2(serviceName string) Generic {
	_ = "STUB: not implemented"
	return *new(Generic)
}

func BinaryPbGeneric(svcName, packageName string) Generic {
	_ = "STUB: not implemented"
	return *new(Generic)
}

func MapThriftGeneric(p DescriptorProvider) (Generic, error) {
	_ = "STUB: not implemented"
	return *new(Generic), nil
}

func MapThriftGenericForJSON(p DescriptorProvider) (Generic, error) {
	_ = "STUB: not implemented"
	return *new(Generic), nil
}

func HTTPThriftGeneric(p DescriptorProvider, opts ...Option) (Generic, error) {
	_ = "STUB: not implemented"
	return *new(Generic), nil
}

func HTTPPbThriftGeneric(p DescriptorProvider, pbp PbDescriptorProvider) (Generic, error) {
	_ = "STUB: not implemented"
	return *new(Generic), nil
}

func JSONThriftGeneric(p DescriptorProvider, opts ...Option) (Generic, error) {
	_ = "STUB: not implemented"
	return *new(Generic), nil
}

func JSONPbGeneric(p PbDescriptorProviderDynamicGo, opts ...Option) (Generic, error) {
	_ = "STUB: not implemented"
	return *new(Generic), nil
}

func SetBinaryWithBase64(g Generic, enable bool) error { _ = "STUB: not implemented"; return nil }

func SetBinaryWithByteSlice(g Generic, enable bool) error { _ = "STUB: not implemented"; return nil }

type SetFieldsForEmptyStructMode uint8

const (
	NotSetFields SetFieldsForEmptyStructMode = iota

	SetNonOptiontionalFields

	SetAllFields
)

func EnableSetFieldsForEmptyStruct(g Generic, mode SetFieldsForEmptyStructMode) error {
	_ = "STUB: not implemented"
	return nil
}

var thriftCodec = thrift.NewThriftCodec()

var pbCodec = protobuf.NewProtobufCodec()

type binaryThriftGeneric struct{}

func (g *binaryThriftGeneric) PayloadCodecType() serviceinfo.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(serviceinfo.PayloadCodec)
}

func (g *binaryThriftGeneric) PayloadCodec() remote.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(remote.PayloadCodec)
}

func (g *binaryThriftGeneric) GenericMethod() serviceinfo.GenericMethodFunc {
	_ = "STUB: not implemented"
	return *new(serviceinfo.GenericMethodFunc)
}

func (g *binaryThriftGeneric) GetExtra(key string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (g *binaryThriftGeneric) Close() error { _ = "STUB: not implemented"; return nil }

func (g *binaryThriftGeneric) IDLServiceName() string { _ = "STUB: not implemented"; return "" }

type binaryThriftGenericV2 struct {
	codec *binaryThriftCodecV2
}

func (b *binaryThriftGenericV2) IDLServiceName() string { _ = "STUB: not implemented"; return "" }

func (b *binaryThriftGenericV2) PayloadCodecType() serviceinfo.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(serviceinfo.PayloadCodec)
}

func (b *binaryThriftGenericV2) GenericMethod() serviceinfo.GenericMethodFunc {
	_ = "STUB: not implemented"
	return *new(serviceinfo.GenericMethodFunc)
}

func (b *binaryThriftGenericV2) GetExtra(key string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (b *binaryThriftGenericV2) Close() error { _ = "STUB: not implemented"; return nil }

type binaryPbGeneric struct {
	codec *binaryPbCodec
}

func (g *binaryPbGeneric) PayloadCodecType() serviceinfo.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(serviceinfo.PayloadCodec)
}

func (g *binaryPbGeneric) GenericMethod() serviceinfo.GenericMethodFunc {
	_ = "STUB: not implemented"
	return *new(serviceinfo.GenericMethodFunc)
}

func (g *binaryPbGeneric) Close() error { _ = "STUB: not implemented"; return nil }

func (g *binaryPbGeneric) IDLServiceName() string { _ = "STUB: not implemented"; return "" }

func (g *binaryPbGeneric) GetExtra(key string) interface{} { _ = "STUB: not implemented"; return nil }

type mapThriftGeneric struct {
	codec *mapThriftCodec
}

func (g *mapThriftGeneric) PayloadCodecType() serviceinfo.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(serviceinfo.PayloadCodec)
}

func (g *mapThriftGeneric) GenericMethod() serviceinfo.GenericMethodFunc {
	_ = "STUB: not implemented"
	return *new(serviceinfo.GenericMethodFunc)
}

func (g *mapThriftGeneric) Close() error { _ = "STUB: not implemented"; return nil }

func (g *mapThriftGeneric) IDLServiceName() string { _ = "STUB: not implemented"; return "" }

func (g *mapThriftGeneric) GetExtra(key string) interface{} { _ = "STUB: not implemented"; return nil }

type jsonThriftGeneric struct {
	codec *jsonThriftCodec
}

func (g *jsonThriftGeneric) PayloadCodecType() serviceinfo.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(serviceinfo.PayloadCodec)
}

func (g *jsonThriftGeneric) GenericMethod() serviceinfo.GenericMethodFunc {
	_ = "STUB: not implemented"
	return *new(serviceinfo.GenericMethodFunc)
}

func (g *jsonThriftGeneric) Close() error { _ = "STUB: not implemented"; return nil }

func (g *jsonThriftGeneric) IDLServiceName() string { _ = "STUB: not implemented"; return "" }

func (g *jsonThriftGeneric) GetExtra(key string) interface{} { _ = "STUB: not implemented"; return nil }

type jsonPbGeneric struct {
	codec *jsonPbCodec
}

func (g *jsonPbGeneric) PayloadCodecType() serviceinfo.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(serviceinfo.PayloadCodec)
}

func (g *jsonPbGeneric) GenericMethod() serviceinfo.GenericMethodFunc {
	_ = "STUB: not implemented"
	return *new(serviceinfo.GenericMethodFunc)
}

func (g *jsonPbGeneric) Close() error { _ = "STUB: not implemented"; return nil }

func (g *jsonPbGeneric) IDLServiceName() string { _ = "STUB: not implemented"; return "" }

func (g *jsonPbGeneric) GetExtra(key string) interface{} { _ = "STUB: not implemented"; return nil }

type httpThriftGeneric struct {
	codec *httpThriftCodec
}

func (g *httpThriftGeneric) PayloadCodecType() serviceinfo.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(serviceinfo.PayloadCodec)
}

func (g *httpThriftGeneric) GenericMethod() serviceinfo.GenericMethodFunc {
	_ = "STUB: not implemented"
	return *new(serviceinfo.GenericMethodFunc)
}

func (g *httpThriftGeneric) Close() error { _ = "STUB: not implemented"; return nil }

func (g *httpThriftGeneric) IDLServiceName() string { _ = "STUB: not implemented"; return "" }

func (g *httpThriftGeneric) GetExtra(key string) interface{} { _ = "STUB: not implemented"; return nil }

type httpPbThriftGeneric struct {
	codec *httpPbThriftCodec
}

func (g *httpPbThriftGeneric) PayloadCodecType() serviceinfo.PayloadCodec {
	_ = "STUB: not implemented"
	return *new(serviceinfo.PayloadCodec)
}

func (g *httpPbThriftGeneric) GenericMethod() serviceinfo.GenericMethodFunc {
	_ = "STUB: not implemented"
	return *new(serviceinfo.GenericMethodFunc)
}

func (g *httpPbThriftGeneric) Close() error { _ = "STUB: not implemented"; return nil }

func (g *httpPbThriftGeneric) IDLServiceName() string { _ = "STUB: not implemented"; return "" }

func (g *httpPbThriftGeneric) GetExtra(key string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func newMethodsFunc(readWriterGetter messageReaderWriterGetter) func(sm Method) serviceinfo.MethodInfo {
	_ = "STUB: not implemented"
	return nil
}

func newMethodsMap(readWriterGetter messageReaderWriterGetter) map[serviceinfo.StreamingMode]serviceinfo.MethodInfo {
	_ = "STUB: not implemented"
	return nil
}

func newMethodInfo(readWriterGetter messageReaderWriterGetter, sm serviceinfo.StreamingMode, oneway bool) serviceinfo.MethodInfo {
	_ = "STUB: not implemented"
	return *new(serviceinfo.MethodInfo)
}

type methodInfo struct {
	serviceinfo.MethodInfo
	oneway bool
}

func (m *methodInfo) OneWay() bool { _ = "STUB: not implemented"; return false }
