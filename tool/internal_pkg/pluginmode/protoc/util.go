package protoc

func goSanitized(s string) string { _ = "STUB: not implemented"; return "" }

type pathElements struct {
	module string
	prefix string
}

func (p *pathElements) getImportPath(pkg string) (path string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}
