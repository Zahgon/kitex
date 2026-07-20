package descriptor

import (
	"sync"
)

type Router interface {
	Handle(rt Route)

	Lookup(req *HTTPRequest) (*FunctionDescriptor, error)
}

type router struct {
	trees      map[string]*node
	maxParams  uint16
	paramsPool sync.Pool
}

func NewRouter() Router { _ = "STUB: not implemented"; return *new(Router) }

func (r *router) getParams() *Params { _ = "STUB: not implemented"; return nil }

func (r *router) putParams(ps *Params) { _ = "STUB: not implemented"; return }

func (r *router) Handle(rt Route) { _ = "STUB: not implemented"; return }

func (r *router) Lookup(req *HTTPRequest) (*FunctionDescriptor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
