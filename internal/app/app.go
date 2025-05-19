package app

import (
	"context"
	"fmt"
	"log/slog"

	"highload-architect/internal/app/container"

	"github.com/meshapi/go-shutdown"

	httpserver "highload-architect/internal/app/http-server"
	"highload-architect/internal/config"
)

type Runner func() error

type App struct {
	ctx     context.Context
	cfg     *config.Config
	runners map[string]Runner
	logger  *slog.Logger
}

func New(ctx context.Context, cfg *config.Config, logger *slog.Logger) (*App, error) {
	cont, err := container.New(cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("init container: %w", err)
	}

	server := httpserver.New(&cfg.HTTPServer, cont, logger)
	logger.Info(fmt.Sprintf("Init Http-Server (%s:%d)", cfg.HTTPServer.Host, cfg.HTTPServer.Port))

	// runners - добавление процессов которые необходимо запустить
	runners := make(map[string]Runner)
	runners["http-server"] = server.Start

	// shutdown - добавление коллбеков, необходимых для GC
	shutdown.AddSequence(shutdown.HandlerFuncWithName("container", cont.Shutdown))
	shutdown.AddSequence(shutdown.HandlerFuncWithName("http-server", server.Shutdown))

	return &App{
		ctx:     ctx,
		cfg:     cfg,
		runners: runners,
		logger:  logger,
	}, nil
}
