package errs

import (
	"errors"
)

var (
	ErrInvalidInput    = errors.New("invalid input")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")
	ErrNotFound        = errors.New("not found")
	ErrConflict        = errors.New("conflict")
	ErrPayloadTooLarge = errors.New("payload too large")
	ErrTooManyRequests = errors.New("too many requests")
	ErrInternal        = errors.New("internal server error")
	ErrUnavailable     = errors.New("service unavailable")
)

type wrappedError struct {
	cause error
	msg   string
}

func (e *wrappedError) Error() string {
	return e.msg
}

func (e *wrappedError) Unwrap() error {
	return e.cause
}

// Wrap attaches msg to cause. A nil cause becomes ErrInternal.
func Wrap(cause error, msg string) error {
	if cause == nil {
		cause = ErrInternal
	}

	return &wrappedError{cause: cause, msg: msg}
}

// Message returns the message set by Wrap, or "" if err has none.
func Message(err error) string {
	if e, ok := errors.AsType[*wrappedError](err); ok {
		return e.msg
	}

	return ""
}
