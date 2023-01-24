package pureerror

import (
	"errors"
	"fmt"
)

type PureError interface {
	error
	Code() string
	Is(error) bool
	Unwrap() error
	Why(string) PureError
	Whyf(string, ...any) PureError
}

type pureError struct {
	code    string
	msg     string
	wrapped error
}

func (e *pureError) Code() string {
	if e == nil {
		return ""
	}
	return e.code
}

func (e *pureError) Error() string {
	if e == nil {
		return ""
	}
	s := e.code
	if e.msg != "" {
		s = fmt.Sprintf("%s: %s", s, e.msg)
	}
	if e.wrapped != nil {
		s = fmt.Sprintf("%s: %s", s, e.wrapped.Error())
	}
	return s
}

func (e *pureError) Is(f error) bool {
	var err *pureError
	if !errors.As(f, &err) {
		return false
	}
	return e.Code() == err.Code()
}

func (e *pureError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.wrapped
}

func (e *pureError) Why(s string) PureError {
	if e == nil {
		return e
	}
	e.msg = s
	return e
}

func (e *pureError) Whyf(format string, a ...any) PureError {
	return e.Why(fmt.Sprintf(format, a...))
}

func New(code string, err error) PureError {
	return &pureError{
		code:    code,
		wrapped: err,
	}
}
