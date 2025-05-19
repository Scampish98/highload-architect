package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"highload-architect/pkg/httputil"
)

func healthCheck(db *sqlx.DB) httputil.HandlerFunc {
	return func(c *gin.Context) error {
		err := db.Ping()
		if err != nil {
			return fmt.Errorf("ping db: %w", err)
		}

		c.JSON(http.StatusOK, "ok")
		return nil
	}
}
