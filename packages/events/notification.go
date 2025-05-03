package events

import (
	"errors"
	"fmt"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type NotificationClient interface {
	SendSMS(phone string, message string) error
}

type notificationClient struct {
	config config.AppConfig
}

func NewNotificationClient(config config.AppConfig) NotificationClient {
	return &notificationClient{
		config: config,
	}
}

func (c notificationClient) SendSMS(phone string, message string) error {
	accountSid := c.config.TwilioAccountSid
	authToken := c.config.TwilioAuthToken
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phone)
	params.SetFrom(c.config.TwilioFromPhoneNumber)
	params.SetBody(message)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error Sending SMS: " + err.Error())
		return errors.New("failed to send sms")
	}
	return nil
}
