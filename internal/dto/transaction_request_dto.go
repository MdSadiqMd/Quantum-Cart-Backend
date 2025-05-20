package dto

import "github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"

type SellerOrderDetails struct {
	Name            string   `json:"name"`
	ImageUrls       []string `json:"image_urls"`
	Price           string   `json:"price"`
	Quantity        uint     `json:"quantity"`
	CustomerName    string   `json:"customer_name"`
	CustomerEmail   string   `json:"customer_email"`
	CustomerPhone   string   `json:"customer_phone"`
	CustomerAddress string   `json:"customer_address"`
	ProductId       uint     `json:"product_id"`
	OrderRefNumber  string   `json:"order_ref_number"`
	OrderStatus     int      `json:"order_status"`
	OrderItemId     uint     `json:"order_item_id"`
	CreatedAt       string   `json:"created_at"`
}

type CreatePaymentRequest struct {
	UserId       uint                 `json:"user_id"`
	OrderId      string               `json:"order_id"`
	PaymentId    string               `json:"payment_id"`
	Status       models.PaymentStatus `json:"status"`
	ClientSecret string               `json:"client"`
	Amount       float64              `json:"amount"`
}
