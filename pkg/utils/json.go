package utils

import (
	"strings"
	"unicode/utf8"
)

const (
	EmptyJSON  = "{}"
	Comma      = ','
	Colon      = ':'
	DQuotation = '"'
	LeftBrace  = '{'
	RightBrace = '}'
)

const (
	t1 = 0x00
	tx = 0x80
	t2 = 0xC0
	t3 = 0xE0
	t4 = 0xF0
	t5 = 0xF8

	maskx = 0x3F

	rune1Max = 1<<7 - 1
	rune2Max = 1<<11 - 1
	rune3Max = 1<<16 - 1

	surrogateMin = 0xD800
	surrogateMax = 0xDFFF

	maxRune   = '\U0010FFFF'
	runeError = '\uFFFD'

	hex = "0123456789abcdef"
)

func _Map2JSONStr(mapInfo map[string]string) (str string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func _JSONStr2Map(jsonStr string) (mapInfo map[string]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readString(buf []byte, idx, lastIdx int) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func readByte(buf []byte, idx, lastIdx int) (byte, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func nextToken(buf []byte, idx, lastIdx int) (byte, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func checkNull(c byte, data []byte, idx, lastIdx int) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func readU4(buf []byte, idx, lastIdx int) (rune, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func readEscapedChar(c byte, buf []byte, idx int, str []byte, lastIdx int) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func wrapStrWithQuotation(s string, strBuilder *strings.Builder) { _ = "STUB: not implemented"; return }

func appendRune(p []byte, r rune) []byte { _ = "STUB: not implemented"; return nil }

var htmlSafeSet = [utf8.RuneSelf]bool{
	' ':      true,
	'!':      true,
	'"':      false,
	'#':      true,
	'$':      true,
	'%':      true,
	'&':      false,
	'\'':     true,
	'(':      true,
	')':      true,
	'*':      true,
	'+':      true,
	',':      true,
	'-':      true,
	'.':      true,
	'/':      true,
	'0':      true,
	'1':      true,
	'2':      true,
	'3':      true,
	'4':      true,
	'5':      true,
	'6':      true,
	'7':      true,
	'8':      true,
	'9':      true,
	':':      true,
	';':      true,
	'<':      false,
	'=':      true,
	'>':      false,
	'?':      true,
	'@':      true,
	'A':      true,
	'B':      true,
	'C':      true,
	'D':      true,
	'E':      true,
	'F':      true,
	'G':      true,
	'H':      true,
	'I':      true,
	'J':      true,
	'K':      true,
	'L':      true,
	'M':      true,
	'N':      true,
	'O':      true,
	'P':      true,
	'Q':      true,
	'R':      true,
	'S':      true,
	'T':      true,
	'U':      true,
	'V':      true,
	'W':      true,
	'X':      true,
	'Y':      true,
	'Z':      true,
	'[':      true,
	'\\':     false,
	']':      true,
	'^':      true,
	'_':      true,
	'`':      true,
	'a':      true,
	'b':      true,
	'c':      true,
	'd':      true,
	'e':      true,
	'f':      true,
	'g':      true,
	'h':      true,
	'i':      true,
	'j':      true,
	'k':      true,
	'l':      true,
	'm':      true,
	'n':      true,
	'o':      true,
	'p':      true,
	'q':      true,
	'r':      true,
	's':      true,
	't':      true,
	'u':      true,
	'v':      true,
	'w':      true,
	'x':      true,
	'y':      true,
	'z':      true,
	'{':      true,
	'|':      true,
	'}':      true,
	'~':      true,
	'\u007f': true,
}
