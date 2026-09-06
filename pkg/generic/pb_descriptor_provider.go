package generic

import (
	dproto "github.com/cloudwego/dynamicgo/proto"

	"github.com/cloudwego/kitex/pkg/generic/proto"
)

type PbDescriptorProvider interface {
	Closer

	Provide() <-chan proto.ServiceDescriptor
}

type PbDescriptorProviderDynamicGo interface {
	Closer

	Provide() <-chan *dproto.ServiceDescriptor
}
