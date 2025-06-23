package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
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

func (r UserDBRepo) Search(ctx context.Context, filter entities.UserFilter) ([]*entities.User, error) {
	query := fmt.Sprintf(`SELECT %s FROM users WHERE 1 = 1`, fieldsStr)
	params := make(map[string]any)

	if filter.FirstNameLike != "" {
		query += ` AND first_name LIKE :first_name_part `
		params["first_name_part"] = fmt.Sprintf("%s%%", filter.FirstNameLike)
	}

	if filter.LastNameLike != "" {
		query += ` AND last_name LIKE :last_name_part `
		params["last_name_part"] = fmt.Sprintf("%s%%", filter.LastNameLike)
	}

	query += ` ORDER BY id`

	stmt, err := r.conn.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("create query: %w", err)
	}
	defer func(ctx context.Context, stmt *sqlx.NamedStmt) {
		if err := stmt.Close(); err != nil {
			r.logger.ErrorContext(ctx, "failed to close insert statement")
		}
	}(ctx, stmt)

	r.logger.DebugContext(ctx, "try exec query",
		slog.String("query", query),
		slog.String("stmt.QueryString", stmt.QueryString),
		slog.String("stmt.Params", fmt.Sprintf("%+v", stmt.Params)),
		slog.String("params", fmt.Sprintf("%+v", params)),
	)

	var users []dbUser

	err = stmt.SelectContext(ctx, &users, params)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return convertManyToEntity(users), nil
}

func (r UserDBRepo) GetByID(ctx context.Context, userID entities.UserID) (*entities.User, error) {
	var row dbUser

	query := fmt.Sprintf(`SELECT %s FROM users WHERE id = $1`, fieldsStr)

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

	query := fmt.Sprintf(`SELECT %s FROM users WHERE username = $1`, fieldsStr)

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
