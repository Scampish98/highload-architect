package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"highload-architect/internal/apperror"

	"github.com/jmoiron/sqlx"

	"highload-architect/internal/entities"
)

type UserDBRepo struct {
	conn   *sqlx.DB
	logger *slog.Logger
}

func NewDBRepo(conn *sqlx.DB, logger *slog.Logger) *UserDBRepo {
	return &UserDBRepo{
		conn:   conn,
		logger: logger,
	}
}

func (r UserDBRepo) GetByID(ctx context.Context, userID entities.UserID) (*entities.User, error) {
	var row dbUser

	query := fmt.Sprintf(`SELECT %s FROM users WHERE id = $1`, strings.Join(fields, ","))

	err := r.conn.GetContext(ctx, &row, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrUserNotFound
		}

		return nil, fmt.Errorf("exec query: %w", err)
	}

	return convertToEntity(row), nil
}

func (r UserDBRepo) GetByUsername(ctx context.Context, username entities.Username) (*entities.User, error) {
	var row dbUser

	query := fmt.Sprintf(`SELECT %s FROM users WHERE username = $1`, strings.Join(fields, ","))

	r.logger.DebugContext(ctx, "try get user by username",
		slog.String("query", query),
		slog.String("username", string(username)),
	)

	err := r.conn.GetContext(ctx, &row, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrUserNotFound
		}

		return nil, fmt.Errorf("exec query: %w", err)
	}

	return convertToEntity(row), nil
}

func (r UserDBRepo) Create(ctx context.Context, user *entities.User) error {
	query := `
INSERT INTO users (username, password, first_name, last_name, sex, birthdate, biography, city, created_at, updated_at) 
VALUES (:username, :password, :first_name, :last_name, :sex, :birthdate, :biography, :city, NOW(), NOW()) 
RETURNING id`

	stmt, err := r.conn.PrepareNamedContext(ctx, query)
	if err != nil {
		return fmt.Errorf("create query: %w", err)
	}
	defer func(ctx context.Context, stmt *sqlx.NamedStmt) {
		if err := stmt.Close(); err != nil {
			r.logger.ErrorContext(ctx, "failed to close insert statement")
		}
	}(ctx, stmt)

	var id uint64

	var birthdate *time.Time
	if user.Birthdate != nil {
		bdate := time.Time(*user.Birthdate)
		birthdate = &bdate
	}

	params := map[string]any{
		"username":   user.Username,
		"password":   user.Password,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"sex":        user.Sex,
		"birthdate":  birthdate,
		"biography":  user.Biography,
		"city":       user.City,
	}

	r.logger.DebugContext(ctx, "try exec query",
		slog.String("query", stmt.QueryString),
		slog.String("params", fmt.Sprintf("%+v", params)),
	)

	err = stmt.QueryRowxContext(ctx, params).Scan(&id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	user.ID = entities.UserID(id)
	return nil
}
