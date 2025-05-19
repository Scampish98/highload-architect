package config

import (
	"fmt"
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

func New() (*Config, error) {
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

	defaults.SetDefaults(cfg)
	return cfg, nil
}
