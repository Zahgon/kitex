package server

import (
	"context"
	"sync"

	igeneric "github.com/cloudwego/kitex/internal/generic"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

var notAllowBinaryGenericCtx = igeneric.WithGenericStreamingMode(context.Background(), serviceinfo.StreamingMode(-1))

type service struct {
	svcInfo              *serviceinfo.ServiceInfo
	handler              interface{}
	unknownMethodHandler interface{}
}

func newService(svcInfo *serviceinfo.ServiceInfo, handler interface{}) *service {
	_ = "STUB: not implemented"
	return nil
}

func (s *service) getHandler(methodName string) interface{} { _ = "STUB: not implemented"; return nil }

type unknownService struct {
	mutex   sync.RWMutex
	svcs    map[string]*service
	handler interface{}
}

func (u *unknownService) getSvc(svcName string) *service { _ = "STUB: not implemented"; return nil }

func (u *unknownService) getOrStoreSvc(svcName string, codecType serviceinfo.PayloadCodec) *service {
	_ = "STUB: not implemented"
	return nil
}

type services struct {
	knownSvcMap map[string]*service

	fallbackSvc     *service
	nonFallbackSvcs []*service

	unknownSvc *unknownService

	binaryThriftGenericV1SvcInfo *serviceinfo.ServiceInfo

	combineSvcInfo *serviceinfo.ServiceInfo

	refuseTrafficWithoutServiceName bool
}

func newServices() *services { _ = "STUB: not implemented"; return nil }

func (s *services) addService(svcInfo *serviceinfo.ServiceInfo, handler interface{}, registerOpts *RegisterOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *services) getKnownSvcInfoMap() map[string]*serviceinfo.ServiceInfo {
	_ = "STUB: not implemented"
	return nil
}

func (s *services) check(refuseTrafficWithoutServiceName bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *services) getService(svcName string) *service { _ = "STUB: not implemented"; return nil }

func (s *services) searchByMethodName(methodName string) *serviceinfo.ServiceInfo {
	_ = "STUB: not implemented"
	return nil
}

func (s *services) SearchService(svcName, methodName string, strict bool, codecType serviceinfo.PayloadCodec) *serviceinfo.ServiceInfo {
	_ = "STUB: not implemented"
	return nil
}

func (s *services) searchUniqueByMethodName(methodName string) (unique *serviceinfo.ServiceInfo) {
	_ = "STUB: not implemented"
	return nil
}

func (s *services) getTargetSvcInfo() *serviceinfo.ServiceInfo {
	_ = "STUB: not implemented"
	return nil
}

func registerBinaryGenericMethodFunc(svcInfo *serviceinfo.ServiceInfo) *serviceinfo.ServiceInfo {
	_ = "STUB: not implemented"
	return nil
}
