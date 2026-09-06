package utils

import (
	"github.com/bytedance/sonic"
)

var sonicConfig = sonic.Config{
	EscapeHTML:     true,
	ValidateString: true,
}.Froze()

func Map2JSONStr(mapInfo map[string]string) (str string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func JSONStr2Map(jsonStr string) (mapInfo map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
