package generator

var DefaultDelimiters = [2]string{"{{", "}}"}

type updateType string

const (
	skip              updateType = "skip"
	cover             updateType = "cover"
	incrementalUpdate updateType = "append"
)

type Update struct {
	Type string `yaml:"type,omitempty"`

	Key string `yaml:"key,omitempty"`

	AppendTpl string `yaml:"append_tpl,omitempty"`

	ImportTpl []string `yaml:"import_tpl,omitempty"`
}

type Template struct {
	Path string `yaml:"path,omitempty"`

	Body string `yaml:"body,omitempty"`

	UpdateBehavior *Update `yaml:"update_behavior,omitempty"`

	LoopMethod bool `yaml:"loop_method,omitempty"`

	LoopService bool `yaml:"loop_service,omitempty"`
}

type customGenerator struct {
	fs       []*File
	pkg      *PackageInfo
	basePath string
}

func NewCustomGenerator(pkg *PackageInfo, basePath string) *customGenerator {
	_ = "STUB: not implemented"
	return nil
}

func (c *customGenerator) loopGenerate(tpl *Template) error { _ = "STUB: not implemented"; return nil }

func (c *customGenerator) commonGenerate(tpl *Template) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *generator) GenerateCustomPackage(pkg *PackageInfo) (fs []*File, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func renderFile(pkg *PackageInfo, outputPath string, tpl *Template) (fs []*File, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readTemplates(dir string) ([]*Template, error) { _ = "STUB: not implemented"; return nil, nil }
