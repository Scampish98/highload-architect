package user

import (
	"highload-architect/internal/transport/http/transformers"
	"net/http"

	"github.com/gin-gonic/gin"

	"highload-architect/internal/entities"
	"highload-architect/pkg/httputil"
)

// Get by ID user
//
//	@Summary		Get user info by ID
//	@Description	Получение информации о пользователе по его идентификатору
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	models.Error
//	@Failure		401	{object}	models.Error
//	@Failure		500	{object}	models.Error
//	@Router			/user/get/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) error {
	id, err := httputil.ParamUint[entities.UserID](c, "id", httputil.Required)
	if err != nil {
		return err
	}

	usr, err := h.userService.GetByID(c.Request.Context(), id)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, transformers.TransformUser(usr))
	return nil
}
