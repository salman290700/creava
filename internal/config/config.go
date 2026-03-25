package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DBUrlMigration string
	SecretJwt      string

	Db_host     string
	Db_user     string
	Db_password string
	Db_port     string
	Db_name     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to load env")
	}
	log.Println("Config Load")
	return &Config{
		Port:           os.Getenv("PORT"),
		DBUrlMigration: os.Getenv("DATABASE_URL"),
		SecretJwt:      os.Getenv("SECRET_JWT"),
		Db_host:        os.Getenv("DB_HOST"),
		Db_user:        os.Getenv("DB_USER"),
		Db_password:    os.Getenv("DB_PASSWORD"),
		Db_port:        os.Getenv("DB_PORT"),
		Db_name:        os.Getenv("DB_NAME"),
	}, nil
}
