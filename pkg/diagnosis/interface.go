package diagnosis

type ProbeName string

type ProbeFunc func() interface{}

type Service interface {
	RegisterProbeFunc(ProbeName, ProbeFunc)
}

func RegisterProbeFunc(svc Service, name ProbeName, pf ProbeFunc) {
	_ = "STUB: not implemented"
	return
}

const (
	ChangeEventsKey    ProbeName = "events"
	ServiceInfosKey    ProbeName = "service_infos"
	FallbackServiceKey ProbeName = "fallback_service"
	UnknownServiceKey  ProbeName = "unknown_service"
	OptionsKey         ProbeName = "options"

	DestServiceKey ProbeName = "dest_service"
	ConnPoolKey    ProbeName = "conn_pool"
	RetryPolicyKey ProbeName = "retry_policy"
)

func WrapAsProbeFunc(data interface{}) ProbeFunc { _ = "STUB: not implemented"; return *new(ProbeFunc) }

var NoopService Service = &noopService{}

type noopService struct{}

func (n noopService) RegisterProbeFunc(name ProbeName, probeFunc ProbeFunc) {
	_ = "STUB: not implemented"
	return
}
