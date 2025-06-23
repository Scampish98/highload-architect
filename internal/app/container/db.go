package container

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"highload-architect/internal/config"

	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Database) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.Type, _getDSNFromConfig(cfg))
	if err != nil {
		return nil, fmt.Errorf("open connection: %w", err)
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return db, nil
}

func _getDSNFromConfig(cfg *config.Database) string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Type,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLmode,
	)
}
