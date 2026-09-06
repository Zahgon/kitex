package descriptor

import (
	"os"

	dthrift "github.com/cloudwego/dynamicgo/thrift"

	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

var isGoTagAliasDisabled = os.Getenv("KITEX_GENERIC_GOTAG_ALIAS_DISABLED") == "True"

type FieldDescriptor struct {
	Name         string
	Alias        string
	ID           int32
	Required     bool
	Optional     bool
	DefaultValue interface{}
	IsException  bool
	Type         *TypeDescriptor
	HTTPMapping  HTTPMapping
	ValueMapping ValueMapping
	GoTagOpt     *GoTagOption
}

type GoTagOption struct {
	IsGoAliasDisabled bool
}

func (d *FieldDescriptor) FieldName() string { _ = "STUB: not implemented"; return "" }

type TypeDescriptor struct {
	Name          string
	Type          Type
	Key           *TypeDescriptor
	Elem          *TypeDescriptor
	Struct        *StructDescriptor
	IsRequestBase bool
}

type StructDescriptor struct {
	Name           string
	FieldsByID     map[int32]*FieldDescriptor
	FieldsByName   map[string]*FieldDescriptor
	RequiredFields map[int32]*FieldDescriptor
	DefaultFields  map[string]*FieldDescriptor
}

func (d *StructDescriptor) CheckRequired(rw map[int32]struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

type FunctionDescriptor struct {
	Name              string
	Oneway            bool
	Request           *TypeDescriptor
	Response          *TypeDescriptor
	HasRequestBase    bool
	IsWithoutWrapping bool
	StreamingMode     serviceinfo.StreamingMode
}

type ServiceDescriptor struct {
	Name               string
	Functions          map[string]*FunctionDescriptor
	Router             Router
	DynamicGoDsc       *dthrift.ServiceDescriptor
	IsCombinedServices bool
}

func (s *ServiceDescriptor) LookupFunctionByMethod(method string) (*FunctionDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
