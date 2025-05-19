package config

import "time"

type HTTPServer struct {
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port" default:"8080"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout" default:"1s"`
	WriteTimeout time.Duration `mapstructure:"write_timeout" default:"1s"`
}
