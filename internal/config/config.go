package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rohitpandeydev/quiz/pkg/logger"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadConfig(log *logger.Logger) (*DBConfig, error) {
	log.Debug("Loading environment variables")
	if err := godotenv.Load(); err != nil {
		log.Error("Error loading .env file: %v", err)
		return nil, err
	}

	config := &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	log.Info("Configuration loaded successfully")
	return config, nil
}
