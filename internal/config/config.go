package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
)

type Config struct {
	Database   Database   `mapstructure:"db"`
	HTTPServer HTTPServer `mapstructure:"http_server"`
	Shutdown   Shutdown   `mapstructure:"shutdown"`

	Auth Auth `mapstructure:"auth"`
}

func New(logger *slog.Logger) (*Config, error) {
	path := os.Getenv("CONFIG")

	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read the configuration file: %w", err)
	}

	logger.Debug("before", slog.String("cfg", fmt.Sprintf("%+v", cfg.Auth.TTL)))
	defaults.SetDefaults(cfg) // TODO dont work
	logger.Debug("after", slog.String("cfg", fmt.Sprintf("%+v", cfg.Auth.TTL)))
	return cfg, nil
}
