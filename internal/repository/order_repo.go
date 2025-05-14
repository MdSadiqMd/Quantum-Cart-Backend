package repository

import (
	"errors"
	"log"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	CreateCartOrders(orders []*models.Order) ([]*models.Order, error)
	GetOrders(userId uint) ([]*models.Order, error)
	GetOrderById(id uint, userId uint) (*models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r orderRepository) CreateOrder(order *models.Order) (*models.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		log.Printf("error in creating order: %v", err)
		return nil, errors.New("failed to create order")
	}
	return order, nil
}

func (r orderRepository) CreateCartOrders(orders []*models.Order) ([]*models.Order, error) {
	err := r.db.Create(&orders).Error
	if err != nil {
		log.Printf("error in creating orders: %v", err)
		return nil, errors.New("failed to create orders")
	}
	return orders, nil
}

func (r orderRepository) GetOrders(userId uint) ([]*models.Order, error) {
	var orders []*models.Order
	err := r.db.Where("user_id = ?", userId).Find(&orders).Error
	if err != nil {
		log.Printf("error in getting orders: %v", err)
		return nil, errors.New("failed to get orders")
	}
	return orders, nil
}

func (r orderRepository) GetOrderById(id uint, userId uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Items").Where("id=? AND user_id=?", id, userId).First(&order).Error
	if err != nil {
		log.Printf("error in getting order: %v", err)
		return nil, errors.New("failed to get order")
	}
	return &order, nil
}
