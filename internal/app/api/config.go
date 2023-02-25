package api

// General instance for API server of REST application
type Config struct {
	// Port
	BindAddr string `toml:"bind_addr"`

	// Logger LeveL
	LoggerLevel string `toml:"logger_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LoggerLevel: "debug",
	}
}
