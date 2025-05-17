package payment

import (
	"errors"
	"fmt"
	"log"

	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)

type PaymentClient interface {
	CreatePayment(amount float64, userId uint, orderId uint) (*stripe.CheckoutSession, error)
	GetPaymentStatus(paymentId string) (*stripe.CheckoutSession, error)
}

type payment struct {
	stripeSecretKey string
	successUrl      string
	cancelURL       string
}

func NewPaymentClient(stripeSecretKey, successUrl, cancelURL string) PaymentClient {
	return &payment{
		stripeSecretKey: stripeSecretKey,
		successUrl:      successUrl,
		cancelURL:       cancelURL,
	}
}

func (p *payment) CreatePayment(amount float64, userId uint, orderId uint) (*stripe.CheckoutSession, error) {
	stripe.Key = p.stripeSecretKey
	amountInCents := amount * 100
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					UnitAmount: stripe.Int64(int64(amountInCents)),
					Currency:   stripe.String(string(stripe.CurrencyINR)),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Clothing"),
					},
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(p.successUrl),
		CancelURL:  stripe.String(p.cancelURL),
	}

	params.AddMetadata("order_id", fmt.Sprintf("%d", orderId))
	params.AddMetadata("user_id", fmt.Sprintf("%d", userId))

	session, err := session.New(params)
	if err != nil {
		log.Println("error in creating stripe session: ", err)
		return nil, errors.New("error in creating stripe session")
	}
	return session, nil
}

func (p *payment) GetPaymentStatus(paymentId string) (*stripe.CheckoutSession, error) {
	stripe.Key = p.stripeSecretKey
	session, err := session.Get(paymentId, nil)
	if err != nil {
		log.Println("error in getting stripe session: ", err)
		return nil, errors.New("error in getting stripe session")
	}
	return session, nil
}
