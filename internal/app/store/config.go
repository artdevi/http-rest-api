package store

type Config struct {
	DatabaseURL string `toml:"database_URL"`
}

func NewConfig() *Config {
	return &Config{}
}
