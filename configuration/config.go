package configuration

import "time"

// Config - project's configuration
type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	PostgreSQL PostgreSQLConfig `mapstructure:"postgresql"`
}

// Serverconfig - project's server's configuration
type ServerConfig struct {
	Host              string        `mapstructure:"host"`
	Port              int           `mapstructure:"port"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout"`
	ReadTimeout       time.Duration `mapstructure:"read_timeout"`
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
}

// Database - project's database's configuration
type PostgreSQLConfig struct {
	Name     string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}
