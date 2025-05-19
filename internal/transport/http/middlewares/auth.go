package middlewares

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"highload-architect/pkg/httputil"

	"highload-architect/internal/app/container"
	"highload-architect/internal/apperror"
	"highload-architect/internal/entities"
)

func Auth(container *container.Container) httputil.HandlerFunc {
	return func(c *gin.Context) error {
		container.Logger.InfoContext(c.Request.Context(), "start auth")
		token := c.Request.Header.Get("x-api-key")

		container.Logger.InfoContext(c.Request.Context(), "token", slog.String("token", token))
		if token == "" {
			return apperror.ErrInvalidToken.WithDetails("token required")
		}

		usr, err := container.Auth.Auth(c.Request.Context(), entities.Token(token))
		if err != nil {
			return err
		}

		c.Set("user", usr)
		c.Next()

		return nil
	}
}
