package user

import (
	"context"
	"fmt"
	"log/slog"

	"highload-architect/internal/entities"
	"highload-architect/internal/infra/user"
)

type UserService struct {
	repo   user.UserRepo
	logger *slog.Logger
}

func New(repo user.UserRepo, logger *slog.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (u *UserService) Register(ctx context.Context, user *entities.User) error {
	err := u.repo.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}

func (u *UserService) GetByUsername(ctx context.Context, username entities.Username) (*entities.User, error) {
	usr, err := u.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("get user by username: %w", err)
	}

	return usr, nil
}

func (u *UserService) GetByID(ctx context.Context, userID entities.UserID) (*entities.User, error) {
	usr, err := u.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user by ID: %w", err)
	}

	return usr, nil
}

func (u *UserService) Search(ctx context.Context, filter entities.UserFilter) ([]*entities.User, error) {
	usrs, err := u.repo.Search(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("search users: %w", err)
	}

	return usrs, nil
}
