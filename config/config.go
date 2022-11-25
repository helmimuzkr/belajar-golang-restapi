package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App struct {
		Port    string
		BaseURL string
	}

	Database struct {
		Host     string
		Port     string
		Username string
		Password string
		DBName   string
	}

	ApiKey struct {
		ApiKeyCategory string
	}
}

func Get() *AppConfig {

	appConfig := initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	appConfig := new(AppConfig)

	// Load file .env di root menggunakan library godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appConfig.App.Port = os.Getenv("APP_PORT")
	appConfig.App.BaseURL = os.Getenv("BASE_URL") + ":" + appConfig.App.Port

	appConfig.Database.Host = os.Getenv("DB_HOST")
	appConfig.Database.Port = os.Getenv("DB_PORT")
	appConfig.Database.Username = os.Getenv("DB_USERNAME")
	appConfig.Database.Password = os.Getenv("DB_PASSWORD")
	appConfig.Database.DBName = os.Getenv("DB_NAME")

	appConfig.ApiKey.ApiKeyCategory = os.Getenv("API_KEY")

	return appConfig
}
