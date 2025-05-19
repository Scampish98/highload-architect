package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/meshapi/go-shutdown"

	"highload-architect/internal/app"
	shutdownlogger "highload-architect/internal/app/shutdown-logger"
	"highload-architect/internal/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	shutdown.SetLogger(shutdownlogger.New(logger))
	shutdown.SetTimeout(cfg.Shutdown.Timeout)
	shutdown.HandlerFuncWithName("cancel context", func() { cancel() })

	application, err := app.New(ctx, cfg, logger)
	if err != nil {
		panic(err)
	}

	application.Start()
}
