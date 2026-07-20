package http

const (
	tcp  = "tcp"
	tcp4 = "tcp4"
	tcp6 = "tcp6"
)

type Resolver interface {
	Resolve(string) (string, error)
}

type ResolverOption func(cfg *resolverConfig)

func WithIPv4() ResolverOption { _ = "STUB: not implemented"; return *new(ResolverOption) }

func WithIPv6() ResolverOption { _ = "STUB: not implemented"; return *new(ResolverOption) }

type resolverConfig struct {
	network string
}

type defaultResolver struct {
	*resolverConfig
}

func NewDefaultResolver(options ...ResolverOption) Resolver {
	_ = "STUB: not implemented"
	return *new(Resolver)
}

func (p *defaultResolver) Resolve(URL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
