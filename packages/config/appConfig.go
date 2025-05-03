package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort            string
	DataSourceName        string
	AppSecret             string
	TwilioAccountSid      string
	TwilioAuthToken       string
	TwilioFromPhoneNumber string
}

func SetupEnv() (cfg AppConfig, err error) {
	godotenv.Load()

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		return AppConfig{}, errors.New("PORT variable not found")
	}

	DataSourceName := os.Getenv("DB_URL")
	if DataSourceName == "" {
		return AppConfig{}, errors.New("DB_URL variable not found")
	}

	appSecret := os.Getenv("APP_SECRET")
	if appSecret == "" {
		return AppConfig{}, errors.New("APP_SECRET variable not found")
	}

	twilioSid := os.Getenv("TWILIO_SID")
	if twilioSid == "" {
		return AppConfig{}, errors.New("TWILIO_SID variable not found")
	}

	twilioToken := os.Getenv("TWILIO_AUTH_TOKEN")
	if twilioToken == "" {
		return AppConfig{}, errors.New("TWILIO_AUTH_TOKEN variable not found")
	}

	twilioFromPhoneNumber := os.Getenv("TWILIO_FROM_NUMBER")
	if twilioFromPhoneNumber == "" {
		return AppConfig{}, errors.New("TWILIO_FROM_NUMBER variable not found")
	}

	return AppConfig{
		ServerPort:            httpPort,
		DataSourceName:        DataSourceName,
		AppSecret:             appSecret,
		TwilioAccountSid:      twilioSid,
		TwilioAuthToken:       twilioToken,
		TwilioFromPhoneNumber: twilioFromPhoneNumber,
	}, nil
}
