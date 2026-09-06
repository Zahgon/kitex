package proto

import (
	"github.com/jhump/protoreflect/desc"
)

type (
	ServiceDescriptor = *desc.ServiceDescriptor
	MessageDescriptor = *desc.MessageDescriptor
)

type Message interface {
	Marshal() ([]byte, error)
	TryGetFieldByNumber(fieldNumber int) (interface{}, error)
	TrySetFieldByNumber(fieldNumber int, val interface{}) error
}

func NewMessage(descriptor MessageDescriptor) Message {
	_ = "STUB: not implemented"
	return *new(Message)
}
