package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"highload-architect/internal/transport/http/models/requests"
	"highload-architect/internal/transport/http/transformers"
	pkgapperror "highload-architect/pkg/apperror"
	pkgstrings "highload-architect/pkg/strings"
	"log/slog"
	"net/http"

	"highload-architect/internal/entities"
)

// Search users
//
//	@Summary		Search users by filter
//	@Description	Получения списка пользователей, удовлетворяющих фильтрам
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			_	query		requests.UserSearchRequest	false	"comment"
//	@Success		200	{object}	models.Users
//	@Failure		400	{object}	models.Error
//	@Failure		401	{object}	models.Error
//	@Failure		500	{object}	models.Error
//	@Router			/user/search/ [get]
func (h *UserHandler) Search(c *gin.Context) error {
	var query requests.UserSearchRequest

	if err := c.ShouldBindQuery(&query); err != nil {
		return pkgapperror.ErrIncorrectParameter.WithInternal(err).WithDetails("invalid query")
	}

	h.logger.Info("search query", slog.String("query", fmt.Sprintf("%+v", query)), slog.String("query string", c.Query("first_name")))

	usrs, err := h.userService.Search(c.Request.Context(), entities.UserFilter{
		FirstNameLike: pkgstrings.Capitalize(query.FirstName),
		LastNameLike:  pkgstrings.Capitalize(query.LastName),
	})

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, transformers.TransformUsers(usrs))
	return nil
}
