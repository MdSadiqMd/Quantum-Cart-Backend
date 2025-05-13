package models

import "time"

type Address struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	AddressLine1 string    `json:"address_line1"`
	AddressLine2 string    `json:"address_line2"`
	City         string    `json:"city"`
	PostCode     uint      `json:"postCode"`
	Country      string    `json:"country"`
	UserId       uint      `json:"user_id"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
	VoidedOn     time.Time `gorm:"default:NULL"`
}
