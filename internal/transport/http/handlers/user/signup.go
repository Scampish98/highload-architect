package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"highload-architect/internal/apperror"

	"highload-architect/internal/entities"
	"highload-architect/internal/transport/http/models/requests"
	pkgapperror "highload-architect/pkg/apperror"
)

// SignUp user
//
//	@Summary		User sign up
//	@Description	Регистрация нового пользователя
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		requests.SignUpRequest	true	"user info"
//	@Success		200		{boolean}	bool
//	@Failure		401		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/user/signup [post]
func (h *UserHandler) SignUp(c *gin.Context) error {
	var reqBody requests.SignUpRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		return pkgapperror.ErrIncorrectParameter.WithInternal(err).WithDetails("invalid request body")
	}

	if err := reqBody.Validate(); err != nil {
		return err
	}

	username := reqBody.ConvertUsername()
	_, err := h.userService.GetByUsername(c.Request.Context(), username)
	if err != nil && !errors.Is(err, apperror.ErrUserNotFound) {
		return err
	} else if err == nil {
		return pkgapperror.ErrIncorrectParameter.WithDetails("user already exists")
	}

	password, err := reqBody.ConvertPassword()
	if err != nil {
		return err
	}

	usr := &entities.User{
		Username:  username,
		Password:  password,
		FirstName: reqBody.ConvertFirstName(),
		LastName:  reqBody.ConvertLastName(),
		Birthdate: reqBody.ConvertBirthdate(),
		Sex:       reqBody.ConvertSex(),
		Biography: reqBody.ConvertBiography(),
		City:      reqBody.ConvertCity(),
	}

	if err := h.userService.Register(c.Request.Context(), usr); err != nil {
		return err
	}

	c.JSON(http.StatusOK, true)
	return nil
}
