package args

import "errors"

type deprecatedFlag struct{}

var errFlagDeprecated = errors.New(`flag is deprecated`)

func (deprecatedFlag) Set(_ string) error { _ = "STUB: not implemented"; return nil }

func (deprecatedFlag) String() string { _ = "STUB: not implemented"; return "" }
