package generator

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

var errNoNewMethod = fmt.Errorf("no new method")

type completer struct {
	allMethods  []*MethodInfo
	handlerPath string
	serviceName string
	streamx     bool
}

func newCompleter(allMethods []*MethodInfo, handlerPath, serviceName string, streamx bool) *completer {
	_ = "STUB: not implemented"
	return nil
}

func parseFuncDecl(fd *ast.FuncDecl) (recvName, funcName string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (c *completer) compare(pkg *ast.Package) []*MethodInfo { _ = "STUB: not implemented"; return nil }

func (c *completer) addImplementations(w io.Writer, newMethods []*MethodInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *completer) addImport(w io.Writer, newMethods []*MethodInfo, fset *token.FileSet, handlerAST *ast.File) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *completer) process(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (c *completer) CompleteMethods() (*File, error) { _ = "STUB: not implemented"; return nil, nil }

type commonCompleter struct {
	path   string
	pkg    *PackageInfo
	update *Update
}

func (c *commonCompleter) Complete() (*File, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *commonCompleter) compare() ([]*MethodInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *commonCompleter) addImport(w io.Writer, newMethods []*MethodInfo, fset *token.FileSet, handlerAST *ast.File) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *commonCompleter) addImplementations(w io.Writer, newMethods []*MethodInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *commonCompleter) parseImports(content string) (imports [][2]string) {
	_ = "STUB: not implemented"
	return nil
}
