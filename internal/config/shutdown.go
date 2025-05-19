package config

import "time"

type Shutdown struct {
	Timeout time.Duration `mapstructure:"timeout" default:"10s"`
}
