package apperror

import (
	"errors"
	"fmt"
)

const (
	UnknownErrorCode        = 0
	InternalServerErrorCode = 1

	NotFoundErrorCode           = 1000
	BadCredentialsErrorCode     = 1001
	IncorrectParameterErrorCode = 1002
)

var (
	ErrUnknown        = New(UnknownErrorCode, "unknown")
	ErrInternalServer = New(InternalServerErrorCode, "internal server error")

	ErrNotFound           = New(NotFoundErrorCode, "not found")
	ErrBadCredentials     = New(BadCredentialsErrorCode, "bad credentials")
	ErrIncorrectParameter = New(IncorrectParameterErrorCode, "incorrect parameter")
)

type AppError struct {
	code     int
	msg      string
	internal error
	details  string
}

func New(code int, msg string) *AppError {
	return &AppError{code: code, msg: msg}
}

func (e AppError) WithInternal(err error) AppError {
	e.internal = err
	return e
}

func (e AppError) WithInternalRecursive(err error) AppError {
	var intErr AppError
	if errors.As(e.internal, &intErr) {
		return intErr.WithInternalRecursive(err)
	}

	e.internal = err
	return e
}

func (e AppError) WithDetails(details string) AppError {
	e.details = details
	return e
}

func (e AppError) Error() string {
	if e.internal != nil {
		return fmt.Sprintf("%s: %s", e.Message(), e.internal.Error())
	}

	return e.Message()
}

func (e AppError) Code() int {
	return e.code
}

func (e AppError) InternalCode() int {
	var intErr AppError
	if e.internal != nil && errors.As(e.internal, &intErr) {
		return intErr.InternalCode()
	}

	return e.code
}

func (e AppError) Message() string {
	if e.details == "" {
		return e.msg
	}

	return fmt.Sprintf("%s (%s)", e.msg, e.details)
}

func (e AppError) Internal() error {
	return e.internal
}
