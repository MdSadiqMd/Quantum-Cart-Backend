package models

import "time"

type Category struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	Name         string    `json:"name" gorm:"index;not null"`
	ParentId     uint      `json:"parent_id"`
	ImageUrl     string    `json:"image_url"`
	Products     []Product `json:"products"`
	DisplayOrder int       `json:"display_order"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"default:current_timestamp"`
	VoidedOn     time.Time `json:"voided_on" gorm:"default:NULL"`
}
