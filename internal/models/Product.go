package models

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"index;not null"`
	Description string    `json:"description"`
	CategoryId  uint      `json:"category_id"`
	ImageUrls   []string  `json:"image_urls" gorm:"type:text[]"`
	Price       float64   `json:"price"`
	UserId      int       `json:"user_id"`
	Stock       uint      `json:"stock"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	VoidedOn    time.Time `json:"voided_on" gorm:"default:NULL"`
}
