package dto

type CreateProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CategoryId  uint     `json:"category_id"`
	ImageUrls   []string `json:"image_urls"`
	Price       float64  `json:"price"`
	UserId      int      `json:"user_id"`
	Stock       uint     `json:"stock"`
}

type UpdateStockRequest struct {
	Stock uint `json:"stock"`
}
