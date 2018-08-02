package config

import (
	"os"
)

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
	SslMode  string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     getEnv("PGHOST", "localhost"),
			Port:     getEnv("PGPORT", "5432"),
			Username: getEnv("PGUSER", "postgres"),
			Password: os.Getenv("PGPASSWORD"),
			DbName:   getEnv("PGDATABASE", "todoapp"),
			SslMode:  getEnv("PGSSLMODE", "require"),
		},
	}
}
