package requests

import (
	"highload-architect/internal/entities"
	"highload-architect/internal/transport/http/models"
	pkgapperror "highload-architect/pkg/apperror"
)

type SignInRequest struct {
	Username string          `json:"username" binding:"required" minLength:"1" maxLength:"50" example:"myusername"`
	Password models.Password `json:"password" binding:"required" minLength:"1" maxLength:"50" example:"1234567"`
}

func (r SignInRequest) Validate() error {
	if r.Username == "" {
		return pkgapperror.ErrIncorrectParameter.WithDetails("username is empty")
	}

	if r.Password == "" {
		return pkgapperror.ErrIncorrectParameter.WithDetails("password is empty")
	}

	return nil
}

func (r SignInRequest) ConvertUsername() entities.Username {
	return entities.Username(r.Username)
}
