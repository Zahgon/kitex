package circuitbreak

import (
	"sync"

	"github.com/bytedance/gopkg/cloud/circuitbreaker"
	"github.com/bytedance/gopkg/collection/skipmap"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/endpoint/cep"
	"github.com/cloudwego/kitex/pkg/event"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

const (
	serviceCBKey  = "service"
	instanceCBKey = "instance"
	cbConfig      = "cb_config"
)

var defaultCBConfig = &CBConfig{Enable: true, ErrRate: 0.5, MinSample: 200}

func GetDefaultCBConfig() CBConfig { _ = "STUB: not implemented"; return *new(CBConfig) }

type CBConfig struct {
	Enable    bool    `json:"enable"`
	ErrRate   float64 `json:"err_rate"`
	MinSample int64   `json:"min_sample"`
}

func (c *CBConfig) DeepCopy() *CBConfig { _ = "STUB: not implemented"; return nil }

func (c *CBConfig) Equals(other *CBConfig) bool { _ = "STUB: not implemented"; return false }

type GenServiceCBKeyFunc func(ri rpcinfo.RPCInfo) string

type instanceCBConfig struct {
	CBConfig
	sync.RWMutex
}

type CBSuite struct {
	servicePanel    circuitbreaker.Panel
	serviceControl  *Control
	instancePanel   circuitbreaker.Panel
	instanceControl *Control

	genServiceCBKey GenServiceCBKeyFunc
	serviceCBConfig *skipmap.StringMap

	instanceCBConfig instanceCBConfig

	events event.Queue

	config CBSuiteConfig
}

func NewCBSuite(genKey GenServiceCBKeyFunc, options ...CBSuiteOption) *CBSuite {
	_ = "STUB: not implemented"
	return nil
}

func (s *CBSuite) ServiceCBMW() endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

func (s *CBSuite) StreamingServiceCBMW() cep.StreamMiddleware {
	_ = "STUB: not implemented"
	return *new(cep.StreamMiddleware)
}

func (s *CBSuite) InstanceCBMW() endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

func (s *CBSuite) ServicePanel() circuitbreaker.Panel {
	_ = "STUB: not implemented"
	return *new(circuitbreaker.Panel)
}

func (s *CBSuite) ServiceControl() *Control { _ = "STUB: not implemented"; return nil }

func (s *CBSuite) UpdateServiceCBConfig(key string, cfg CBConfig) {
	_ = "STUB: not implemented"
	return
}

func (s *CBSuite) UpdateInstanceCBConfig(cfg CBConfig) { _ = "STUB: not implemented"; return }

func (s *CBSuite) SetEventBusAndQueue(bus event.Bus, events event.Queue) {
	_ = "STUB: not implemented"
	return
}

func (s *CBSuite) Dump() interface{} { _ = "STUB: not implemented"; return nil }

func (s *CBSuite) Close() error { _ = "STUB: not implemented"; return nil }

func (s *CBSuite) initServiceCB() { _ = "STUB: not implemented"; return }

func (s *CBSuite) initInstanceCB() { _ = "STUB: not implemented"; return }

func (s *CBSuite) onStateChange(level, key string, oldState, newState circuitbreaker.State, m circuitbreaker.Metricer) {
	_ = "STUB: not implemented"
	return
}

func (s *CBSuite) onServiceStateChange(key string, oldState, newState circuitbreaker.State, m circuitbreaker.Metricer) {
	_ = "STUB: not implemented"
	return
}

func (s *CBSuite) onInstanceStateChange(key string, oldState, newState circuitbreaker.State, m circuitbreaker.Metricer) {
	_ = "STUB: not implemented"
	return
}

func (s *CBSuite) discoveryChangeHandler(e *event.Event) { _ = "STUB: not implemented"; return }

func (s *CBSuite) svcTripFunc(key string) circuitbreaker.TripFunc {
	_ = "STUB: not implemented"
	return *new(circuitbreaker.TripFunc)
}

func (s *CBSuite) insTripFunc(key string) circuitbreaker.TripFunc {
	_ = "STUB: not implemented"
	return *new(circuitbreaker.TripFunc)
}

func cbDebugInfo(panel circuitbreaker.Panel) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (s *CBSuite) configInfo() map[string]interface{} { _ = "STUB: not implemented"; return nil }

func RPCInfo2Key(ri rpcinfo.RPCInfo) string { _ = "STUB: not implemented"; return "" }
