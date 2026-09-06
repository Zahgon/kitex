package descriptor

import (
	"reflect"
	"regexp"
)

var escape = regexp.MustCompile(`\\.`)

type FiledMapping interface {
	Handle(field *FieldDescriptor)
}

type NewFieldMapping func(value string) FiledMapping

var GoTagAnnatition = NewBAMAnnotation("go.tag", NewGoTag)

type goTag struct {
	tag reflect.StructTag
}

var NewGoTag NewFieldMapping = func(value string) FiledMapping {
	value = escape.ReplaceAllStringFunc(value, func(m string) string {
		if m[1] == '"' {
			return m[1:]
		}
		return m
	})
	return &goTag{reflect.StructTag(value)}
}

func (m *goTag) Handle(field *FieldDescriptor) { _ = "STUB: not implemented"; return }
