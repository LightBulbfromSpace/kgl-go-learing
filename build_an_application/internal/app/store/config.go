package store

import "fmt"

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "username"
	password = "password"
	dbname   = "players"
)

func NewConfig() *Config {
	return &Config{}
}

func NewConfigForTest() *Config {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return &Config{psqlInfo}
}
