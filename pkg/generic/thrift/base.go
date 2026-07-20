package thrift

import (
	"github.com/cloudwego/gopkg/protocol/thrift/base"
)

type Base = base.Base

func NewBase() *Base { _ = "STUB: not implemented"; return nil }

type BaseResp = base.BaseResp

func NewBaseResp() *BaseResp { _ = "STUB: not implemented"; return nil }

func MergeBase(jsonBase, frameworkBase Base) Base { _ = "STUB: not implemented"; return *new(Base) }

func mergeBaseAny(jsonBase any, frkBase *Base) *Base { _ = "STUB: not implemented"; return nil }
