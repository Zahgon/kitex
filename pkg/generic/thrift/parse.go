package thrift

import (
	"github.com/cloudwego/thriftgo/generator/golang/streaming"
	"github.com/cloudwego/thriftgo/parser"

	"github.com/cloudwego/kitex/pkg/generic/descriptor"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

const (
	initRecursionDepth = 0
)

type ParseMode int

const (
	LastServiceOnly ParseMode = iota

	FirstServiceOnly

	CombineServices
)

var defaultParseMode = LastServiceOnly

func DefaultParseMode() ParseMode { _ = "STUB: not implemented"; return *new(ParseMode) }

func SetDefaultParseMode(m ParseMode) { _ = "STUB: not implemented"; return }

func Parse(tree *parser.Thrift, mode ParseMode, opts ...ParseOption) (*descriptor.ServiceDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTargetService(svcs []*parser.Service, serviceName string) ([]*parser.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type pair struct {
	tree *parser.Thrift
	data interface{}
}

func getAllSvcs(svc *parser.Service, tree *parser.Thrift, visitedSvcs map[*parser.Service]bool) chan *pair {
	_ = "STUB: not implemented"
	return nil
}

func addFunction(fn *parser.Function, tree *parser.Thrift, sDsc *descriptor.ServiceDescriptor, structsCache map[string]*descriptor.TypeDescriptor, opts *parseOptions) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func streamingMode(st *streaming.Streaming) serviceinfo.StreamingMode {
	_ = "STUB: not implemented"
	return *new(serviceinfo.StreamingMode)
}

func parseRequest(isStream bool, field *parser.Field, tree *parser.Thrift, structsCache map[string]*descriptor.TypeDescriptor, opts *parseOptions) (req *descriptor.TypeDescriptor, hasRequestBase bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func parseResponse(isStream bool, fn *parser.Function, tree *parser.Thrift, structsCache map[string]*descriptor.TypeDescriptor, opts *parseOptions) (*descriptor.TypeDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var builtinTypes = map[string]*descriptor.TypeDescriptor{
	"void":   {Name: "void", Type: descriptor.VOID, Struct: new(descriptor.StructDescriptor)},
	"bool":   {Name: "bool", Type: descriptor.BOOL},
	"byte":   {Name: "byte", Type: descriptor.BYTE},
	"i8":     {Name: "i8", Type: descriptor.I08},
	"i16":    {Name: "i16", Type: descriptor.I16},
	"i32":    {Name: "i32", Type: descriptor.I32},
	"i64":    {Name: "i64", Type: descriptor.I64},
	"double": {Name: "double", Type: descriptor.DOUBLE},
	"string": {Name: "string", Type: descriptor.STRING},
	"binary": {Name: "binary", Type: descriptor.STRING},
}

func parseType(t *parser.Type, tree *parser.Thrift, cache map[string]*descriptor.TypeDescriptor, recursionDepth int, opt *parseOptions) (*descriptor.TypeDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parse(tree *parser.Thrift, name string, t *parser.Type, v *parser.ConstValue) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func onBool(tree *parser.Thrift, name string, t *parser.Type, v *parser.ConstValue) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func onInt(tree *parser.Thrift, name string, t *parser.Type, v *parser.ConstValue) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func onDouble(tree *parser.Thrift, name string, t *parser.Type, v *parser.ConstValue) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func onStr(tree *parser.Thrift, name string, t *parser.Type, v *parser.ConstValue) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func onEnum(tree *parser.Thrift, name string, t *parser.Type, v *parser.ConstValue) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func onSetOrList(tree *parser.Thrift, name string, t *parser.Type, v *parser.ConstValue) (res []interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func onMap(tree *parser.Thrift, name string, t *parser.Type, v *parser.ConstValue) (res map[interface{}]interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func onStructLike(tree *parser.Thrift, name string, t *parser.Type, v *parser.ConstValue) (res map[string]interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getStructLike(tree *parser.Thrift, t *parser.Type) (ast *parser.Thrift, s *parser.StructLike, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func getIDValue(tree *parser.Thrift, name string, t *parser.Type, extra *parser.ConstValueExtra) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func bin2str(t *parser.Type) *parser.Type { _ = "STUB: not implemented"; return nil }
