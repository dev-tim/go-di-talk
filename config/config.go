package config

type Config struct {
	ServerPort string
}

func NewConfig() *Config {
	return &Config{
		ServerPort: "8080",
	}
}
