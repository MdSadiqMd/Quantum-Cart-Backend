package repository

import (
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"gorm.io/gorm"
)

type CatalogRepository interface {
	CreateCategory(category *models.Category) error
	FindCategories() ([]*models.Category, error)
	FindCategoryById(id uint) (*models.Category, error)
	EditCategory(category *models.Category) (*models.Category, error)
	DeleteCategory(category *models.Category) error
}

type catalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{
		db: db,
	}
}

func (r catalogRepository) CreateCategory(category *models.Category) error {
	err := r.db.Create(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func (r catalogRepository) FindCategories() ([]*models.Category, error) {
	var categories []*models.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r catalogRepository) FindCategoryById(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r catalogRepository) EditCategory(category *models.Category) (*models.Category, error) {
	err := r.db.Save(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r catalogRepository) DeleteCategory(category *models.Category) error {
	err := r.db.Delete(&category).Error
	if err != nil {
		return err
	}
	return nil
}
