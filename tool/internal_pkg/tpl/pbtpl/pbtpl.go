package pbtpl

import (
	"io"
)

type Args struct {
	Services []*Service
	StreamX  bool
}

type Service struct {
	Name    string
	Methods []*Method
}

type Method struct {
	Name    string
	ReqType string
	ResType string

	ClientStream bool
	ServerStream bool
}

func Render(w io.Writer, in *Args) { _ = "STUB: not implemented"; return }

func renderService(fm func(format string, aa ...any), s *Service, streamx bool) {
	_ = "STUB: not implemented"
	return
}

func notPtr(s string) string { _ = "STUB: not implemented"; return "" }
