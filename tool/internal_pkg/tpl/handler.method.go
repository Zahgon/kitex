package tpl

var HandlerMethodsTpl string = `{{define "HandlerMethod"}}
{{range .AllMethods}}
{{- if or .ClientStreaming .ServerStreaming}}
func (s *{{$.ServiceName}}Impl) {{.Name}}({{if $.StreamX}}ctx context.Context, {{end}}{{if not .ClientStreaming}}{{range .Args}}{{LowerFirst .Name}} {{.Type}}, {{end}}{{end}}stream {{.PkgRefName}}.{{.ServiceName}}_{{.RawName}}Server) (err error) {	
	println("{{.Name}} called")
	return
}
{{- else}}
{{- if .Void}}
// {{.Name}} implements the {{.ServiceName}}Impl interface.
{{- if .Oneway}}
// Oneway methods are not guaranteed to receive 100% of the requests sent by the client.
// And the client may not perceive the loss of requests due to network packet loss.
// If possible, do not use oneway methods.
{{- end}}
func (s *{{$.ServiceName}}Impl) {{.Name}}(ctx context.Context {{- range .Args}}, {{LowerFirst .Name}} {{.Type}}{{end}}) (err error) {
	// TODO: Your code here...
	return
}
{{else -}}
// {{.Name}} implements the {{.ServiceName}}Impl interface.
func (s *{{$.ServiceName}}Impl) {{.Name}}(ctx context.Context {{range .Args}}, {{LowerFirst .Name}} {{.Type}}{{end}} ) (resp {{.Resp.Type}}, err error) {
	// TODO: Your code here...
	return
}
{{end}}
{{end}}
{{end}}
{{end}}{{/* define "HandlerMethod" */}}
`
