package models

import (
	"time"

	"github.com/lib/pq"
)

type OrderStatus string

const (
	OrderStatusDelivered  OrderStatus = "initial"
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "yet_to_deliver"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type Order struct {
	ID             uint        `json:"id" gorm:"primary_key"`
	UserId         uint        `json:"user_id" gorm:"index;not null"`
	Status         OrderStatus `json:"status" gorm:"default:initial"`
	Amount         float64     `json:"amount"`
	TransactionId  string      `json:"transaction_id"`
	OrderRefNumber string      `json:"order_ref_number"`
	PaymentId      string      `json:"payment_id"`
	Items          []OrderItem `json:"items"`
	CreatedAt      time.Time   `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt      time.Time   `json:"updated_at" gorm:"default:current_timestamp"`
	VoidedOn       time.Time   `json:"voided_on" gorm:"default:NULL"`
}

type OrderItem struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	OrderId   uint           `json:"order_id" gorm:"index;not null"`
	ProductId uint           `json:"product_id"`
	Name      string         `json:"name"`
	ImageUrls pq.StringArray `json:"image_urls" gorm:"type:text[]"`
	SellerId  uint           `json:"seller_id"`
	Price     float64        `json:"price"`
	Quantity  uint           `json:"quantity"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:current_timestamp"`
	VoidedOn  time.Time      `json:"voided_on" gorm:"default:NULL"`
}
