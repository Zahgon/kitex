package utils

type KitexArgs interface {
	GetFirstArgument() interface{}
}

type KitexResult interface {
	GetResult() interface{}
	SetSuccess(interface{})
}
