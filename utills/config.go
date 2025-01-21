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
		HOST:     os.Getenv("HOST"),
		GIN_MODE: os.Getenv("GIN_MODE"),
	}

	return nil

}
