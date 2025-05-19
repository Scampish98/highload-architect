package app

import (
	"fmt"

	"github.com/meshapi/go-shutdown"
)

func (a *App) Start() {
	for name, runFn := range a.runners {
		name := name
		runFn := runFn
		go func() {
			if err := runFn(); err != nil {
				a.logger.ErrorContext(a.ctx, fmt.Sprintf("failed to start %s: %v", name, err))
				shutdown.Trigger(a.ctx)
			}
		}()
	}

	shutdown.WaitForInterrupt()
}
