package models

import "time"

type Cart struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	UserId    uint      `json:"user_id"`
	ProductId uint      `json:"product_id"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"image_url"`
	Price     float64   `json:"price"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	VoidedOn  time.Time `json:"voided_on" gorm:"default:NULL"`
}
