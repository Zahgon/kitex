package descriptor

import (
	"net/http"
)

type Renderer interface {
	Render(w http.ResponseWriter, body interface{}) error
	WriteContentType(w http.ResponseWriter)
}

type JsonRenderer struct{}

func (j JsonRenderer) Render(w http.ResponseWriter, body interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (j JsonRenderer) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

type PbRenderer struct{}

func (p PbRenderer) Render(w http.ResponseWriter, body interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (p PbRenderer) WriteContentType(w http.ResponseWriter) { _ = "STUB: not implemented"; return }
