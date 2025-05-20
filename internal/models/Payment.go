package models

import "time"

type PaymentStatus string

const (
	PaymentInit    PaymentStatus = "initial"
	PaymentPending PaymentStatus = "pending"
	PaymentSuccess PaymentStatus = "success"
	PaymentFailed  PaymentStatus = "failed"
)

type Payment struct {
	ID            uint          `json:"id" gorm:"primary_key"`
	UserId        uint          `josn:"user_id"`
	CaptureMethod string        `json:"capture_method"`
	Amount        float64       `json:"amount"`
	OrderId       string        `json:"order_id"`
	CustomerId    string        `json:"customer_id"`
	PaymentId     string        `json:"payment_id"`
	ClientSecret  string        `json:"client_secret"`
	Status        PaymentStatus `json:"status" gorm:"default:initial"`
	Response      string        `json:"response"`
	CreatedAt     time.Time     `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt     time.Time     `json:"updated_at" gorm:"default:current_timestamp"`
	VoidedOn      time.Time     `json:"voided_on" gorm:"default:NULL"`
}
