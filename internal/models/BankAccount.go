package models

import "time"

type BankAccount struct {
	Id          uint      `json:"id" gorm:"primary_key"`
	UserId      uint      `json:"user_id"`
	BankAccount uint      `json:"bank_account" gorm:"index;unique;not null"`
	SwiftCode   uint      `json:"swift_code"`
	PaymentType string    `json:"payment_type"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	VoidedOn    time.Time `json:"voided_on" gorm:"default:NULL"`
}
