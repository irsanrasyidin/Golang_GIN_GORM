package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ApiConfig struct {
	AppPort string
}
type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}
type Config struct {
	ApiConfig
	DbConfig
}

func (c Config) readConfigFile() Config {
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	c.ApiConfig = ApiConfig{
		AppPort: os.Getenv("APP_PORT"),
	}
	return c
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	cfg := Config{}
	return cfg.readConfigFile()
}
