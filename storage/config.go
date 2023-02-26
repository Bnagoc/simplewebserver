package storage

type Config struct {
	// DB connaction line
	DatabaseURI string `toml:"database_uri"`
}

func NewConfig() *Config {
	return &Config{}
}
