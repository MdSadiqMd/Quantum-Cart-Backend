package repository

import (
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	FindProducts(limit, offset int) ([]*models.Product, error)
	FindProductbyId(id uint) (*models.Product, error)
	FindSellerProducts(id uint) ([]*models.Product, error)
	UpdateProduct(id uint, product *models.Product) (*models.Product, error)
	DeleteProduct(product *models.Product) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r productRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	err := r.db.Model(&models.Product{}).Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r productRepository) FindProducts(limit, offset int) ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.Find(&products).Limit(10).Offset(0).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r productRepository) FindProductbyId(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r productRepository) FindSellerProducts(id uint) ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.Where("user_id = ?", id).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r productRepository) UpdateProduct(id uint, product *models.Product) (*models.Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r productRepository) DeleteProduct(product *models.Product) (*models.Product, error) {
	err := r.db.Delete(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
