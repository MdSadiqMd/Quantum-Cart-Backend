package dto

import "github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"

type CreateOrderRequest struct {
	ProductId uint     `json:"product_id"`
	Quantity  uint     `json:"quantity"`
	Name      string   `json:"name"`
	ImageUrls []string `json:"image_urls"`
	SellerId  uint     `json:"seller_id"`
	Price     float64  `json:"price"`
}

type OrderResponse struct {
	Order    *models.Order `json:"order"`
	OrderRef string        `json:"order_ref"`
}

type BulkOrderRequest struct {
	IsBulk bool                  `json:"is_bulk"`
	Orders []*CreateOrderRequest `json:"orders"`
}

type BulkOrderResponse struct {
	Orders    []*models.Order `json:"orders"`
	OrderRefs []string        `json:"order_refs"`
}
