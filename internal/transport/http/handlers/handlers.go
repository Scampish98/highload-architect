package handlers

import (
	"github.com/gin-gonic/gin"

	"highload-architect/pkg/httputil"

	"highload-architect/internal/app/container"
	"highload-architect/internal/transport/http/handlers/user"
)

func RegUnauthHTTPHandlers(r *gin.RouterGroup, c *container.Container) {
	r.GET("/health-check", httputil.Wrap(healthCheck(c.DB)))

	user.RegUnauthHTTPHandlers(r.Group("/user"), c.User, c.Auth, c.Logger)
}

func RegAuthHTTPHandlers(r *gin.RouterGroup, c *container.Container) {
	user.RegAuthHTTPHandlers(r.Group("/user"), c.User, c.Auth, c.Logger)
}
