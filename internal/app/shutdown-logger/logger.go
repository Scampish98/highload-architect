package shutdownlogger

import "log/slog"

type ShutdownLogger struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *ShutdownLogger {
	return &ShutdownLogger{
		logger: logger,
	}
}

func (l *ShutdownLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *ShutdownLogger) Error(msg string) {
	l.logger.Error(msg)
}
