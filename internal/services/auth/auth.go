package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"highload-architect/internal/apperror"
	"highload-architect/internal/config"
	"highload-architect/internal/entities"
	"highload-architect/internal/services/user"
	pkgapperror "highload-architect/pkg/apperror"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	cfg    *config.Auth
	user   *user.UserService
	logger *slog.Logger

	jwtParser *jwt.Parser
}

func New(cfg *config.Auth, user *user.UserService, logger *slog.Logger) *Auth {
	return &Auth{
		cfg:    cfg,
		user:   user,
		logger: logger,

		jwtParser: jwt.NewParser(jwt.WithExpirationRequired()),
	}
}

func (auth *Auth) Auth(ctx context.Context, token entities.Token) (*entities.User, error) {
	payload, err := auth.parseToken(token)
	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}

	usr, err := auth.user.GetByUsername(ctx, payload.username)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (auth *Auth) GenerateToken(user *entities.User) (entities.Token, error) {
	payload := tokenPayload{username: user.Username}

	claims := payload.toJwtClaims()
	claims["exp"] = time.Now().Add(auth.cfg.TTL).Unix()

	auth.logger.Debug("claims", slog.String("values", fmt.Sprintf("%+v", claims)))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(auth.cfg.Secret))
	if err != nil {
		return "", fmt.Errorf("sign token: %w", err)
	}

	return entities.Token(t), nil
}

func (auth *Auth) parseToken(token entities.Token) (tokenPayload, error) {
	var claims jwt.MapClaims

	_, err := auth.jwtParser.ParseWithClaims(string(token), &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.cfg.Secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return tokenPayload{}, apperror.ErrTokenExpired.WithInternalRecursive(err)
		}

		if errors.Is(err, jwt.ErrSignatureInvalid) || errors.Is(err, jwt.ErrTokenMalformed) {
			return tokenPayload{}, apperror.ErrInvalidToken.WithInternalRecursive(err)
		}

		return tokenPayload{}, pkgapperror.ErrUnknown.WithInternal(err)
	}

	payload, err := parseTokenPayload(claims)
	if err != nil {
		return tokenPayload{}, fmt.Errorf("parse token payload: %w", err)
	}

	return payload, nil
}
