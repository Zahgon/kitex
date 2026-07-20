package descriptor

type Annotation interface {
	Equal(key, value string) bool

	Handle() interface{}
}

var annotations = []Annotation{}

func RegisterAnnotation(an Annotation) { _ = "STUB: not implemented"; return }

func init() {

	RegisterAnnotation(APIQueryAnnotation)
	RegisterAnnotation(APIPathAnnotation)
	RegisterAnnotation(APIHeaderAnnotation)
	RegisterAnnotation(APICookieAnnotation)
	RegisterAnnotation(APIBodyAnnotation)
	RegisterAnnotation(APIHttpCodeAnnotation)
	RegisterAnnotation(APINoneAnnotation)
	RegisterAnnotation(APIRawBodyAnnotation)

	RegisterAnnotation(APIGetAnnotation)
	RegisterAnnotation(APIPostAnnotation)
	RegisterAnnotation(APIPutAnnotation)
	RegisterAnnotation(APIDeleteAnnotation)

	RegisterAnnotation(GoTagAnnatition)

	RegisterAnnotation(APIJSConvAnnotation)

	RegisterAnnotation(apiVdAnnotation)
	RegisterAnnotation(apiSerializerAnnotation)
	RegisterAnnotation(apiParamAnnotation)
	RegisterAnnotation(apiBaseURLAnnotation)
	RegisterAnnotation(apiGenPathAnnotation)
	RegisterAnnotation(apiVersionAnnotation)
	RegisterAnnotation(apiTagAnnotation)
	RegisterAnnotation(apiVDAnnotation)
}

func FindAnnotation(key, value string) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type bamAnnotation struct {
	key    string
	handle interface{}
}

func NewBAMAnnotation(key string, handle interface{}) Annotation {
	_ = "STUB: not implemented"
	return *new(Annotation)
}

func (a *bamAnnotation) Equal(key, value string) bool { _ = "STUB: not implemented"; return false }

func (a *bamAnnotation) Handle() interface{} { _ = "STUB: not implemented"; return nil }

type noneAnnotation struct {
	key string
}

func NewNoneAnnotation(key string) Annotation { _ = "STUB: not implemented"; return *new(Annotation) }

func (a *noneAnnotation) Equal(key, value string) bool { _ = "STUB: not implemented"; return false }

func (a *noneAnnotation) Handle() interface{} { _ = "STUB: not implemented"; return nil }

var (
	apiVdAnnotation         = NewNoneAnnotation("api.vd")
	apiSerializerAnnotation = NewNoneAnnotation("api.serializer")
	apiParamAnnotation      = NewNoneAnnotation("api.param")
	apiBaseURLAnnotation    = NewNoneAnnotation("api.baseurl")
	apiGenPathAnnotation    = NewNoneAnnotation("api.gen_path")
	apiVersionAnnotation    = NewNoneAnnotation("api.version")
	apiTagAnnotation        = NewNoneAnnotation("api.tag")
	apiVDAnnotation         = NewNoneAnnotation("api.vd")
)

type noneWithValueAnnotation struct {
	key, value string
}

func NewNoneWithValueAnnotation(key, value string) Annotation {
	_ = "STUB: not implemented"
	return *new(Annotation)
}

func (a *noneWithValueAnnotation) Equal(key, value string) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *noneWithValueAnnotation) Handle() interface{} { _ = "STUB: not implemented"; return nil }
