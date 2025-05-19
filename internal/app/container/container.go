package container

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"

	"highload-architect/internal/config"
	userdb "highload-architect/internal/infra/user"
	authservice "highload-architect/internal/services/auth"
	userservice "highload-architect/internal/services/user"
)

type Container struct {
	DB *sqlx.DB

	Logger *slog.Logger

	Auth *authservice.Auth
	User *userservice.UserService
}

func New(cfg *config.Config, logger *slog.Logger) (*Container, error) {
	logger.Info("db dsn", slog.String("dsn", fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.Type,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SSLmode,
	)))
	db, err := NewDB(&cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	userRepo := userdb.NewDBRepo(db, logger)
	userService := userservice.New(userRepo, logger)

	authService := authservice.New(&cfg.Auth, userService, logger)

	return &Container{
		DB: db,

		Logger: logger,

		Auth: authService,
		User: userService,
	}, nil
}

func (c *Container) Shutdown(ctx context.Context) error {
	if err := c.DB.Close(); err != nil {
		return err
	}

	return nil
}
