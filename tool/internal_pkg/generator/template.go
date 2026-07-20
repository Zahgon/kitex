package generator

type APIExtension struct {
	ImportPaths []string `json:"import_paths,omitempty" yaml:"import_paths,omitempty"`

	ExtendOption string `json:"extend_option,omitempty" yaml:"extend_option,omitempty"`

	ExtendFile string `json:"extend_file,omitempty" yaml:"extend_file,omitempty"`
}

type TemplateExtension struct {
	FeatureNames []string `json:"feature_names,omitempty" yaml:"feature_names,omitempty"`

	EnableFeatures []string `json:"enable_features,omitempty" yaml:"enable_features,omitempty"`

	Dependencies map[string]string `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`

	ExtendClient *APIExtension `json:"extend_client,omitempty" yaml:"extend_client,omitempty"`

	ExtendServer *APIExtension `json:"extend_server,omitempty" yaml:"extend_server,omitempty"`

	ExtendInvoker *APIExtension `json:"extend_invoker,omitempty" yaml:"extend_invoker,omitempty"`
}

func (p *TemplateExtension) FromJSONFile(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TemplateExtension) ToJSONFile(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TemplateExtension) FromYAMLFile(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TemplateExtension) ToYAMLFile(filename string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TemplateExtension) Merge(other *TemplateExtension) { _ = "STUB: not implemented"; return }

func (a *APIExtension) Merge(other *APIExtension) { _ = "STUB: not implemented"; return }
