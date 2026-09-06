package generic

import (
	"context"
	"sync"

	dproto "github.com/cloudwego/dynamicgo/proto"

	"github.com/cloudwego/kitex/pkg/generic/proto"
)

type PbContentProvider struct {
	closeOnce sync.Once
	svcs      chan proto.ServiceDescriptor
}

type PbFileProviderWithDynamicGo struct {
	closeOnce sync.Once
	svcs      chan *dproto.ServiceDescriptor
}

var (
	_ PbDescriptorProvider          = (*PbContentProvider)(nil)
	_ PbDescriptorProviderDynamicGo = (*PbFileProviderWithDynamicGo)(nil)
)

func NewPbContentProvider(main string, includes map[string]string) (PbDescriptorProvider, error) {
	_ = "STUB: not implemented"
	return *new(PbDescriptorProvider), nil
}

func (p *PbContentProvider) UpdateIDL(main string, includes map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseProto(main string, includes map[string]string) (proto.ServiceDescriptor, error) {
	_ = "STUB: not implemented"
	return *new(proto.ServiceDescriptor), nil
}

func (p *PbContentProvider) Provide() <-chan proto.ServiceDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (p *PbContentProvider) Close() error { _ = "STUB: not implemented"; return nil }

func NewPbFileProviderWithDynamicGo(main string, ctx context.Context, options dproto.Options, importDirs ...string) (PbDescriptorProviderDynamicGo, error) {
	_ = "STUB: not implemented"
	return *new(PbDescriptorProviderDynamicGo), nil
}

func NewPbContentProviderWithDynamicGo(ctx context.Context, options dproto.Options, mainPath, mainContent string, includes map[string]string) (PbDescriptorProviderDynamicGo, error) {
	_ = "STUB: not implemented"
	return *new(PbDescriptorProviderDynamicGo), nil
}

func (p *PbFileProviderWithDynamicGo) UpdateIDL(ctx context.Context, options dproto.Options, mainPath, mainContent string, includes map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *PbFileProviderWithDynamicGo) Provide() <-chan *dproto.ServiceDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func (p *PbFileProviderWithDynamicGo) Close() error { _ = "STUB: not implemented"; return nil }
