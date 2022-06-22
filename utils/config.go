package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port           string
	PostgresConfig PostgresConfig
	RedisConfig    RedisConfig
}

type PostgresConfig struct {
	Password string
	Db       string
	User     string
	Port     string
	Host     string
}

type RedisConfig struct {
	Host string
	Port string
}

func GetConfigs() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Using Dev Env: %v", err)
	}
	config := Config{
		Port: fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		PostgresConfig: PostgresConfig{
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Db:       os.Getenv("POSTGRES_DB"),
			User:     os.Getenv("POSTGRES_USER"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Host:     os.Getenv("POSTGRES_HOST"),
		},
		RedisConfig: RedisConfig{
			Host: os.Getenv("REDIS_HOST"),
			Port: os.Getenv("REDIS_PORT"),
		},
	}
	return config
}
