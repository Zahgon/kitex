package descriptor

import (
	"net/http"
	"net/url"

	"github.com/bytedance/sonic/ast"
	dhttp "github.com/cloudwego/dynamicgo/http"
)

type MIMEType string

const (
	MIMEApplicationJson     = "application/json"
	MIMEApplicationProtobuf = "application/x-protobuf"
)

var (
	_ dhttp.RequestGetter  = &HTTPRequest{}
	_ dhttp.ResponseSetter = &HTTPResponse{}
)

type HTTPRequest struct {
	Params      *Params
	Request     *http.Request
	RawBody     []byte
	Body        map[string]interface{}
	GeneralBody interface{}
	ContentType MIMEType
	cookies     map[string]string
	query       url.Values
	bodyMap     *ast.Node
}

func (req *HTTPRequest) GetHeader(key string) string { _ = "STUB: not implemented"; return "" }

func (req *HTTPRequest) GetCookie(key string) string { _ = "STUB: not implemented"; return "" }

func (req *HTTPRequest) GetQuery(key string) string { _ = "STUB: not implemented"; return "" }

func (req *HTTPRequest) GetBody() []byte { _ = "STUB: not implemented"; return nil }

func (req *HTTPRequest) GetMethod() string { _ = "STUB: not implemented"; return "" }

func (req *HTTPRequest) GetPath() string { _ = "STUB: not implemented"; return "" }

func (req *HTTPRequest) GetHost() string { _ = "STUB: not implemented"; return "" }

func (req *HTTPRequest) GetParam(key string) string { _ = "STUB: not implemented"; return "" }

func (req *HTTPRequest) GetMapBody(key string) string { _ = "STUB: not implemented"; return "" }

func (req *HTTPRequest) GetPostForm(key string) string { _ = "STUB: not implemented"; return "" }

func (req *HTTPRequest) GetUri() string { _ = "STUB: not implemented"; return "" }

func (req *HTTPRequest) initializeBodyMap() error { _ = "STUB: not implemented"; return nil }

type HTTPResponse struct {
	Header      http.Header
	StatusCode  int32
	RawBody     []byte
	Body        map[string]interface{}
	GeneralBody interface{}
	ContentType MIMEType
	Renderer    Renderer
}

func NewHTTPResponse() *HTTPResponse { _ = "STUB: not implemented"; return nil }

func (resp *HTTPResponse) SetStatusCode(code int) error { _ = "STUB: not implemented"; return nil }

func (resp *HTTPResponse) SetHeader(key, val string) error { _ = "STUB: not implemented"; return nil }

func (resp *HTTPResponse) SetCookie(key, val string) error { _ = "STUB: not implemented"; return nil }

func (resp *HTTPResponse) SetRawBody(body []byte) error { _ = "STUB: not implemented"; return nil }

func NewHTTPPbResponse(initBody interface{}) *HTTPResponse { _ = "STUB: not implemented"; return nil }

func NewGeneralHTTPResponse(contentType MIMEType, initBody interface{}, renderer Renderer) *HTTPResponse {
	_ = "STUB: not implemented"
	return nil
}

func (resp *HTTPResponse) Write(w http.ResponseWriter) error { _ = "STUB: not implemented"; return nil }

type Param struct {
	Key   string
	Value string
}

type Params struct {
	params   []Param
	recycle  func(*Params)
	recycled bool
}

func (ps *Params) Recycle() { _ = "STUB: not implemented"; return }

func (ps *Params) ByName(name string) string { _ = "STUB: not implemented"; return "" }
