package grpcframe

import (
	"errors"

	"golang.org/x/net/http2"
)

type connError struct {
	Code   http2.ErrCode
	Reason string
}

func (e connError) Error() string { _ = "STUB: not implemented"; return "" }

type pseudoHeaderError string

func (e pseudoHeaderError) Error() string { _ = "STUB: not implemented"; return "" }

type duplicatePseudoHeaderError string

func (e duplicatePseudoHeaderError) Error() string { _ = "STUB: not implemented"; return "" }

type headerFieldNameError string

func (e headerFieldNameError) Error() string { _ = "STUB: not implemented"; return "" }

type headerFieldValueError string

func (e headerFieldValueError) Error() string { _ = "STUB: not implemented"; return "" }

var (
	errMixPseudoHeaderTypes = errors.New("mix of request and response pseudo headers")
	errPseudoAfterRegular   = errors.New("pseudo header field after regular")
)
