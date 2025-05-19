package middlewares

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"highload-architect/internal/transport/http/models"
	pkgapperror "highload-architect/pkg/apperror"

	"highload-architect/internal/app/container"
)

func Error(container *container.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()

		logArgs := []any{
			slog.String("uri", c.Request.RequestURI),
			slog.String("method", c.Request.Method),
			slog.Int("status", c.Writer.Status()),
			slog.Duration("latency", time.Since(startTime)),
		}

		err := c.Errors.Last()
		if err != nil {
			var appErr pkgapperror.AppError
			if errors.As(err, &appErr) {
				c.AbortWithStatusJSON(AppErrorToHTTPErrorCode(appErr), models.Error{
					Code:    appErr.InternalCode(),
					Message: appErr.Error(),
				})

				return
			}

			container.Logger.ErrorContext(
				c.Request.Context(),
				fmt.Sprintf("http request error: %s", err.Error()),
				logArgs...,
			)

			c.AbortWithStatusJSON(http.StatusInternalServerError, models.Error{
				Code:    pkgapperror.UnknownErrorCode,
				Message: err.Error(),
			})
		}
	}
}

func AppErrorToHTTPErrorCode(err pkgapperror.AppError) int {
	switch err.Code() {
	case pkgapperror.BadCredentialsErrorCode:
		return http.StatusUnauthorized
	case pkgapperror.NotFoundErrorCode:
		return http.StatusNotFound
	case pkgapperror.IncorrectParameterErrorCode:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
