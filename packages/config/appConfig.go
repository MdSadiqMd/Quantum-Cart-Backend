package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnv() (cfg AppConfig, err error) {
	godotenv.Load()

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		return AppConfig{}, errors.New("PORT variable not found")
	}

	return AppConfig{ServerPort: httpPort}, nil
}
