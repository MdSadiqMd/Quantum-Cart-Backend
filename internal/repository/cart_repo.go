package repository

import (
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartRepository interface {
	FindCartItems(userId uint) ([]*models.Cart, error)
	FindCartItem(userId uint, productId uint) (*models.Cart, error)
	CreateCart(cart models.Cart) (*models.Cart, error)
	UpdateCart(cart models.Cart) (*models.Cart, error)
	DeleteCartById(id uint) error
	DeleteCartItems(userId uint) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r cartRepository) FindCartItems(userId uint) ([]*models.Cart, error) {
	var carts []*models.Cart
	err := r.db.Where("user_id = ?", userId).Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (r cartRepository) FindCartItem(userId uint, productId uint) (*models.Cart, error) {
	cartItem := &models.Cart{}
	err := r.db.Where("user_id = ? AND product_id = ?", userId, productId).First(&cartItem).Error
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

func (r cartRepository) CreateCart(cart models.Cart) (*models.Cart, error) {
	err := r.db.Create(&cart).Error
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (r cartRepository) UpdateCart(cart models.Cart) (*models.Cart, error) {
	var UpdateCart models.Cart
	err := r.db.Model(&UpdateCart).Clauses(clause.Returning{}).Where("id = ?", cart.ID).Updates(cart).Error
	if err != nil {
		return nil, err
	}
	return &UpdateCart, nil
}

func (r cartRepository) DeleteCartById(id uint) error {
	err := r.db.Delete(&models.Cart{}, id).Error
	return err
}

func (r cartRepository) DeleteCartItems(userId uint) error {
	err := r.db.Where("user_id = ?", userId).Delete(&models.Cart{}).Error
	return err
}
