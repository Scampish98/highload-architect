package apperror

import "highload-architect/pkg/apperror"

const (
	UserNotFoundErrorCode      = 10000
	IncorrectPasswordErrorCode = 10001
	InvalidTokenErrorCode      = 10002
	TokenExpiredErrorCode      = 10003
)

var (
	errUserNotFound      = apperror.New(UserNotFoundErrorCode, "user not found")
	errIncorrectPassword = apperror.New(IncorrectPasswordErrorCode, "incorrect user password")
	errInvalidToken      = apperror.New(InvalidTokenErrorCode, "invalid token")
	errTokenExpired      = apperror.New(TokenExpiredErrorCode, "token is expired")
)

var (
	ErrUserNotFound      = apperror.ErrNotFound.WithInternal(errUserNotFound)
	ErrIncorrectPassword = apperror.ErrBadCredentials.WithInternal(errIncorrectPassword)
	ErrInvalidToken      = apperror.ErrBadCredentials.WithInternal(errInvalidToken)
	ErrTokenExpired      = apperror.ErrBadCredentials.WithInternal(errTokenExpired)
)
