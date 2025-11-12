package config

type Config struct {
	Env string
}

func LoadConfig() (*Config, error) {
	return &Config{
		Env: "dev",
	}, nil
}
