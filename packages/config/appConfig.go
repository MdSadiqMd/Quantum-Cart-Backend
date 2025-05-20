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
	StripeSecret          string
	PublishableKey        string
	SuccessURL            string
	CancelURL             string
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

	StripeSecret := os.Getenv("STRIPE_SECRET")
	if StripeSecret == "" {
		return AppConfig{}, errors.New("STRIPE_SECRET variable not found")
	}

	PublishableKey := os.Getenv("STRIPE_PUBLISHABLE_KEY")
	if PublishableKey == "" {
		return AppConfig{}, errors.New("PUBLISHABLE_KEY variable not found")
	}

	SuccessURL := os.Getenv("SUCCESS_URL")
	if SuccessURL == "" {
		return AppConfig{}, errors.New("SUCCESS_URL variable not found")
	}

	CancelURL := os.Getenv("CANCEL_URL")
	if CancelURL == "" {
		return AppConfig{}, errors.New("CANCEL_URL variable not found")
	}

	return AppConfig{
		ServerPort:            httpPort,
		DataSourceName:        DataSourceName,
		AppSecret:             appSecret,
		TwilioAccountSid:      twilioSid,
		TwilioAuthToken:       twilioToken,
		TwilioFromPhoneNumber: twilioFromPhoneNumber,
		StripeSecret:          StripeSecret,
		PublishableKey:        PublishableKey,
		SuccessURL:            SuccessURL,
		CancelURL:             CancelURL,
	}, nil
}
