package services

import (
	"errors"
	"strings"
	"time"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
)

type CartService struct {
	CartRepo    repository.CartRepository
	ProductRepo repository.ProductRepository
	Auth        helpers.Auth
	Config      config.AppConfig
}

func NewCartService(
	cartRepo repository.CartRepository,
	productRepo repository.ProductRepository,
	auth helpers.Auth,
	config config.AppConfig,
) *CartService {
	return &CartService{
		CartRepo:    cartRepo,
		ProductRepo: productRepo,
		Auth:        auth,
		Config:      config,
	}
}

func (s CartService) FindCart(id uint) ([]*models.Cart, float64, error) {
	cart, err := s.CartRepo.FindCartItems(id)
	if err != nil {
		return nil, 0, errors.New("error on finding cart items")
	}

	var totalAmount float64
	for _, item := range cart {
		totalAmount += item.Price * float64(item.Quantity)
	}
	return cart, totalAmount, err
}

func (s CartService) CreateCart(cart dto.CreateCartRequest, user models.User) ([]*models.Cart, error) {
	existingCart, _ := s.CartRepo.FindCartItem(user.Id, cart.ProductId)
	if existingCart == nil {
		product, _ := s.ProductRepo.FindProductbyId(uint(cart.ProductId))
		if product.ID < 1 {
			return nil, errors.New("product not found to create cart item")
		}

		_, err := s.CartRepo.CreateCart(models.Cart{
			UserId:    user.Id,
			ProductId: product.ID,
			Name:      product.Name,
			ImageUrl:  strings.Join(product.ImageUrls, ","),
			Price:     product.Price,
			Quantity:  cart.Quantity,
			UpdatedAt: time.Now(),
		})
		if err != nil {
			return nil, errors.New("error on creating cart item")
		}
	} else {
		if cart.ProductId == 0 {
			return nil, errors.New("please provide a valid product id")
		}

		if cart.Quantity < 1 {
			err := s.CartRepo.DeleteCartById(existingCart.ID)
			if err != nil {
				return nil, errors.New("error on deleting cart item")
			}
		} else {
			existingCart.Quantity = cart.Quantity
			_, err := s.CartRepo.UpdateCart(*existingCart)
			if err != nil {
				return nil, errors.New("error on updating cart item")
			}
		}
	}
	return s.CartRepo.FindCartItems(user.Id)
}
