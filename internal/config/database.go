package config

import "time"

type Database struct {
	Type            string        `mapstructure:"type" default:"postgres"`
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	Name            string        `mapstructure:"name"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	SSLmode         string        `mapstructure:"sslmode"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" default:"100"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time" default:"60s"`
}
