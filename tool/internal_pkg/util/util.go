package util

import (
	"regexp"
)

type StringSlice []string

func (ss *StringSlice) String() string { _ = "STUB: not implemented"; return "" }

func (ss *StringSlice) Set(value string) error { _ = "STUB: not implemented"; return nil }

func FormatCode(code []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func GetGOPATH() (string, error) { _ = "STUB: not implemented"; return "", nil }

func Exists(path string) bool { _ = "STUB: not implemented"; return false }

func LowerFirst(s string) string { _ = "STUB: not implemented"; return "" }

func ReplaceString(s, old, new string, n int) string { _ = "STUB: not implemented"; return "" }

func SnakeString(s string) string { _ = "STUB: not implemented"; return "" }

func UpperFirst(s string) string { _ = "STUB: not implemented"; return "" }

func NotPtr(s string) string { _ = "STUB: not implemented"; return "" }

func SearchGoMod(cwd string) (moduleName, path string, found bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

func RunGitCommand(gitLink string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func CombineOutputPath(outputPath, ns string) string { _ = "STUB: not implemented"; return "" }

func JoinPath(elem ...string) string { _ = "STUB: not implemented"; return "" }

func DownloadFile(remotePath, localPath string) error { _ = "STUB: not implemented"; return nil }

func IDLName(filename string) string { _ = "STUB: not implemented"; return "" }

type Import struct {
	Alias string
	Path  string
}

func SortImports(imps map[string]string, localPrefix string) (ret []Import) {
	_ = "STUB: not implemented"
	return nil
}

func (i Import) PackageName() string { _ = "STUB: not implemented"; return "" }

func PrintlImports(imports []Import) string { _ = "STUB: not implemented"; return "" }

var packageRE = regexp.MustCompile(`package\s+([a-zA-Z0-9_]+)`)

func TruncateAllFastPBFiles(dir string) { _ = "STUB: not implemented"; return }

func TruncateFastPBFile(fn string) (success bool) { _ = "STUB: not implemented"; return false }
