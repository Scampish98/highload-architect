package user

import (
	"log/slog"

	"highload-architect/internal/services/auth"

	"github.com/gin-gonic/gin"

	"highload-architect/pkg/httputil"

	"highload-architect/internal/services/user"
)

type UserHandler struct {
	userService *user.UserService
	authService *auth.Auth
	logger      *slog.Logger
}

func newUserHandler(userService *user.UserService, authService *auth.Auth, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		userService: userService,
		authService: authService,
		logger:      logger,
	}
}

func RegUnauthHTTPHandlers(r *gin.RouterGroup, userService *user.UserService, authService *auth.Auth, logger *slog.Logger) {
	handler := newUserHandler(userService, authService, logger)

	r.POST("/signup", httputil.Wrap(handler.SignUp))
	r.POST("/signin", httputil.Wrap(handler.SignIn))
}

func RegAuthHTTPHandlers(r *gin.RouterGroup, userService *user.UserService, authService *auth.Auth, logger *slog.Logger) {
	handler := newUserHandler(userService, authService, logger)

	r.GET("/get/:id", httputil.Wrap(handler.GetByID))
	r.GET("/search", httputil.Wrap(handler.Search))
}
