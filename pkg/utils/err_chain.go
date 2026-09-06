package utils

const DefaultErrorSeparator = " | "

type ErrChain struct {
	errs []error
	sep  *string
}

func (e *ErrChain) UseSeparator(sep string) { _ = "STUB: not implemented"; return }

func (e *ErrChain) Append(err error) { _ = "STUB: not implemented"; return }

func (e ErrChain) HasError() bool { _ = "STUB: not implemented"; return false }

func (e ErrChain) Error() string { _ = "STUB: not implemented"; return "" }
