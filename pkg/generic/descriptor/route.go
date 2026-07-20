package descriptor

import "net/http"

type Route interface {
	Method() string
	Path() string
	Function() *FunctionDescriptor
}

type NewRoute func(value string, function *FunctionDescriptor) Route

var (
	APIGetAnnotation = NewBAMAnnotation("api.get", NewAPIGet)

	APIPostAnnotation = NewBAMAnnotation("api.post", NewAPIPost)

	APIPutAnnotation = NewBAMAnnotation("api.put", NewAPIPut)

	APIDeleteAnnotation = NewBAMAnnotation("api.delete", NewAPIDelete)
)

var NewAPIGet NewRoute = func(value string, function *FunctionDescriptor) Route {
	return &apiRoute{http.MethodGet, value, function}
}

var NewAPIPost NewRoute = func(value string, function *FunctionDescriptor) Route {
	return &apiRoute{http.MethodPost, value, function}
}

var NewAPIPut NewRoute = func(value string, function *FunctionDescriptor) Route {
	return &apiRoute{http.MethodPut, value, function}
}

var NewAPIDelete NewRoute = func(value string, function *FunctionDescriptor) Route {
	return &apiRoute{http.MethodDelete, value, function}
}

type apiRoute struct {
	method   string
	value    string
	function *FunctionDescriptor
}

func (r *apiRoute) Method() string { _ = "STUB: not implemented"; return "" }

func (r *apiRoute) Path() string { _ = "STUB: not implemented"; return "" }

func (r *apiRoute) Function() *FunctionDescriptor { _ = "STUB: not implemented"; return nil }
