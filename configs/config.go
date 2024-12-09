package configs

import (
	"github.com/joho/godotenv"
	"high-load-final/pkg/models"
	"os"
)

func InitConfigs() (models.AppConfigs, error) {
	env := os.Getenv("ENV")
	if env == "PROD" {
		err := godotenv.Load(".env.prod")
		if err != nil {
			return models.AppConfigs{}, err
		}
	} else if env == "DEMO" {
		err := godotenv.Load(".env.demo")
		if err != nil {
			return models.AppConfigs{}, err
		}
	} else {
		err := godotenv.Load(".env")
		if err != nil {
			return models.AppConfigs{}, err
		}
	}
	database := models.Postgres{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	logs := models.LogsConfig{
		Path: os.Getenv("LOGS_PATH"),
	}

	swagger := models.SwaggerAuth{
		Username: os.Getenv("SWAGGER_USERNAME"),
		Password: os.Getenv("SWAGGER_PASSWORD"),
	}

	return models.AppConfigs{
		Postgres:    database,
		Logs:        logs,
		SwaggerAuth: swagger,
	}, nil
}
