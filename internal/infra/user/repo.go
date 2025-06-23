package user

import (
	"context"

	"highload-architect/internal/entities"
)

type UserRepo interface {
	GetByID(ctx context.Context, userID entities.UserID) (*entities.User, error)
	GetByUsername(ctx context.Context, username entities.Username) (*entities.User, error)
	Create(ctx context.Context, user *entities.User) error
	Search(ctx context.Context, filter entities.UserFilter) ([]*entities.User, error)
}
