package server

import (
	"context"
	"net"
	"sync"

	"github.com/cloudwego/localsession/backup"

	internal_server "github.com/cloudwego/kitex/internal/server"
	"github.com/cloudwego/kitex/pkg/diagnosis"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/endpoint/sep"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/remotesvr"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/serviceinfo"
)

type Server interface {
	RegisterService(svcInfo *serviceinfo.ServiceInfo, handler interface{}, opts ...RegisterOption) error
	GetServiceInfos() map[string]*serviceinfo.ServiceInfo
	Run() error
	Stop() error
}

type server struct {
	opt  *internal_server.Options
	svcs *services

	eps     endpoint.Endpoint
	svr     remotesvr.Server
	stopped sync.Once
	isInit  bool
	isRun   bool

	sync.Mutex
}

func NewServer(ops ...Option) Server { _ = "STUB: not implemented"; return *new(Server) }

func (s *server) init() {
	if s.isInit {
		return
	}
	s.isInit = true
	ctx := fillContext(s.opt)
	if ds := s.opt.DebugService; ds != nil {
		ds.RegisterProbeFunc(diagnosis.OptionsKey, diagnosis.WrapAsProbeFunc(s.opt.DebugInfo))
		ds.RegisterProbeFunc(diagnosis.ChangeEventsKey, s.opt.Events.Dump)
	}
	backup.Init(s.opt.BackupOpt)

	s.buildInvokeChain(ctx)
}

func fillContext(opt *internal_server.Options) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (s *server) initOrResetRPCInfoFunc() func(rpcinfo.RPCInfo, net.Addr) rpcinfo.RPCInfo {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) buildMiddlewares(ctx context.Context) []endpoint.Middleware {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) buildInvokeChain(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *server) RegisterService(svcInfo *serviceinfo.ServiceInfo, handler interface{}, opts ...RegisterOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) GetServiceInfos() map[string]*serviceinfo.ServiceInfo {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) Run() (err error) { _ = "STUB: not implemented"; return nil }

func (s *server) Stop() (err error) { _ = "STUB: not implemented"; return nil }

func (s *server) buildCoreMiddleware() endpoint.Middleware {
	_ = "STUB: not implemented"
	return *new(endpoint.Middleware)
}

func (s *server) unaryOrStreamEndpoint(ctx context.Context) endpoint.Endpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.Endpoint)
}

func (s *server) invokeHandleEndpoint() endpoint.UnaryEndpoint {
	_ = "STUB: not implemented"
	return *new(endpoint.UnaryEndpoint)
}

func (s *server) streamHandleEndpoint() sep.StreamEndpoint {
	_ = "STUB: not implemented"
	return *new(sep.StreamEndpoint)
}

func (s *server) initBasicRemoteOption() { _ = "STUB: not implemented"; return }

func (s *server) richRemoteOption() { _ = "STUB: not implemented"; return }

func (s *server) addBoundHandlers(opt *remote.ServerOption) { _ = "STUB: not implemented"; return }

func (s *server) buildLimiterWithOpt() (handler remote.InboundHandler) {
	_ = "STUB: not implemented"
	return *new(remote.InboundHandler)
}

func (s *server) check() error { _ = "STUB: not implemented"; return nil }

func doAddBoundHandlerToHead(h remote.BoundHandler, opt *remote.ServerOption) {
	_ = "STUB: not implemented"
	return
}

func doAddBoundHandler(h remote.BoundHandler, opt *remote.ServerOption) {
	_ = "STUB: not implemented"
	return
}

func (s *server) newSvrTransHandler() (handler remote.ServerTransHandler, err error) {
	_ = "STUB: not implemented"
	return *new(remote.ServerTransHandler), nil
}

func (s *server) buildRegistryInfo(lAddr net.Addr) { _ = "STUB: not implemented"; return }

func (s *server) registerDebugInfo() { _ = "STUB: not implemented"; return }

func (s *server) waitExit(errCh chan error) error { _ = "STUB: not implemented"; return nil }
