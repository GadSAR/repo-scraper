package utills

import (
	"os"

	"github.com/GadSAR/repo-scraper/models"
	"github.com/joho/godotenv"
)

var Config *models.Config

func LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	Config = &models.Config{
		HOST:     getEnv("HOST", "localhost:8080"),
		GIN_MODE: getEnv("GIN_MODE", "realase"),
		ZAP_MODE: getEnv("ZAP_MODE", "development"),
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value

}
