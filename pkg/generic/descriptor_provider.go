package generic

import (
	dthrift "github.com/cloudwego/dynamicgo/thrift"

	"github.com/cloudwego/kitex/pkg/generic/descriptor"
)

type DescriptorProvider interface {
	Closer

	Provide() <-chan *descriptor.ServiceDescriptor
}

type GetProviderOption interface {
	Option() ProviderOption
}

type ProviderOption struct {
	DynamicGoEnabled bool

	DynamicGoOptions *dthrift.Options
}
