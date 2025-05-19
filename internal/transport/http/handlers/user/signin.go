package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"highload-architect/internal/transport/http/models/requests"
	"highload-architect/internal/transport/http/models/responses"
	pkgapperror "highload-architect/pkg/apperror"
)

// SignUp user
//
//	@Summary		User sign in
//	@Description	Вход для существующего пользователя
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		requests.SignInRequest	true	"user info"
//	@Success		200		{object}	responses.SignInResponse
//	@Failure		401		{object}	models.Error
//	@Failure		500		{object}	models.Error
//	@Router			/user/signin [post]
func (h *UserHandler) SignIn(c *gin.Context) error {
	var reqBody requests.SignInRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		return pkgapperror.ErrIncorrectParameter.WithInternal(err).WithDetails("invalid request body")
	}

	if err := reqBody.Validate(); err != nil {
		return err
	}

	username := reqBody.ConvertUsername()

	usr, err := h.userService.GetByUsername(c.Request.Context(), username)
	if err != nil {
		return err
	}

	if err := reqBody.Password.Compare(string(usr.Password)); err != nil {
		return err
	}

	token, err := h.authService.GenerateToken(usr)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, responses.SignInResponse{Token: string(token)})
	return nil
}
