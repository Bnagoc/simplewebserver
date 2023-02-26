package api

import "simplewebserver/storage"

// General instance for API server of REST application
type Config struct {
	// Port
	BindAddr string `toml:"bind_addr"`

	// Logger LeveL
	LoggerLevel string `toml:"logger_level"`

	// Storage config
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
