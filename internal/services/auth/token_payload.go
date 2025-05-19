package auth

import (
	"github.com/golang-jwt/jwt/v5"

	"highload-architect/internal/apperror"
	"highload-architect/internal/entities"
)

const (
	_fieldUsername = "username"
)

type tokenPayload struct {
	username entities.Username
}

func parseTokenPayload(claims jwt.MapClaims) (tokenPayload, error) {
	rawUsername, ok := claims[_fieldUsername]
	if !ok {
		return tokenPayload{}, apperror.ErrInvalidToken.WithDetails("username not found")
	}

	username, ok := rawUsername.(string)
	if !ok {
		return tokenPayload{}, apperror.ErrInvalidToken.WithDetails("username is not a string")
	}

	return tokenPayload{username: entities.Username(username)}, nil
}

func (t *tokenPayload) toJwtClaims() jwt.MapClaims {
	return jwt.MapClaims{
		_fieldUsername: t.username,
	}
}
