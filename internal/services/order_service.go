package services

import (
	"errors"

	"fmt"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
	"github.com/lib/pq"
)

type OrderService struct {
	OrderRepo repository.OrderRepository
	CartRepo  repository.CartRepository
	Auth      helpers.Auth
	Config    config.AppConfig
}

func (s *OrderService) CreateOrder(userId uint, orders []*dto.CreateOrderRequest, isBulk bool) ([]*models.Order, []string, error) {
	cartItems, err := s.CartRepo.FindCartItems(userId)
	if err != nil {
		return nil, nil, err
	}

	cartMap := make(map[uint]models.Cart)
	for _, item := range cartItems {
		if item.UserId == userId {
			cartMap[item.ProductId] = *item
		}
	}

	var totalAmount float64
	var orderItems []models.OrderItem
	var cartItemsToDelete []uint
	for _, o := range orders {
		if o.ProductId == 0 || o.Quantity <= 0 {
			continue
		}

		totalAmount += o.Price * float64(o.Quantity)
		orderItems = append(orderItems, models.OrderItem{
			ProductId: o.ProductId,
			Quantity:  o.Quantity,
			Name:      o.Name,
			Price:     o.Price,
			ImageUrls: pq.StringArray(o.ImageUrls),
			SellerId:  o.SellerId,
		})

		if cartItem, exists := cartMap[o.ProductId]; exists {
			cartItemsToDelete = append(cartItemsToDelete, cartItem.ID)
		}
	}
	if len(orderItems) == 0 {
		return nil, nil, errors.New("no valid items to order")
	}

	orderRef, _ := helpers.RandomNumbers(8)
	creatingOrder := &models.Order{
		UserId:         userId,
		PaymentId:      "123",
		TransactionId:  "123",
		OrderRefNumber: fmt.Sprint(orderRef),
		Amount:         totalAmount,
		Items:          orderItems,
	}

	var createdOrders []*models.Order
	if isBulk {
		createdOrders, err = s.OrderRepo.CreateCartOrders([]*models.Order{creatingOrder})
	} else {
		var singleOrder *models.Order
		singleOrder, err = s.OrderRepo.CreateOrder(creatingOrder)
		createdOrders = []*models.Order{singleOrder}
	}
	if err != nil {
		return nil, nil, err
	}

	for _, cartID := range cartItemsToDelete {
		if err := s.CartRepo.DeleteCartById(cartID); err != nil {
			return nil, nil, err
		}
	}

	var refs []string
	for _, o := range createdOrders {
		refs = append(refs, o.OrderRefNumber)
	}

	return createdOrders, refs, nil
}

func (s *OrderService) GetOrders(userId uint) ([]*models.Order, error) {
	orders, err := s.OrderRepo.GetOrders(userId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) GetOrderById(id uint, userId uint) (*models.Order, error) {
	order, err := s.OrderRepo.GetOrderById(id, userId)
	if err != nil {
		return nil, err
	}
	return order, nil
}
