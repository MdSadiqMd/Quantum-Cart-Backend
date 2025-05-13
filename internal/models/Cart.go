package models

import "time"

type Cart struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserId    uint      `json:"user_id"`
	ProductId uint      `json:"product_id"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"image_url"`
	SellerId  uint      `json:"seller_id"`
	Price     float64   `json:"price"`
	Quantity  uint      `json:"quantity"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
	VoidedOn  time.Time `gorm:"default:NULL"`
}
