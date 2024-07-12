package config

type Config struct {
	PG   string
	Port string
}

func NewConfig() *Config {
	return &Config{
		PG:   "postgres://db:db@localhost:12323/test_db?sslmode=disable",
		Port: ":12342",
	}
}
