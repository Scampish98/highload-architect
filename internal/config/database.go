package config

type Database struct {
	Type     string `mapstructure:"type" default:"postgres"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	SSLmode  string `mapstructure:"sslmode"`
}
