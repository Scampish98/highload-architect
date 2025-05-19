package config

import "time"

type Auth struct {
	Secret string        `mapstructure:"secret"`
	TTL    time.Duration `mapstructure:"ttl" default:"1h"`
}
