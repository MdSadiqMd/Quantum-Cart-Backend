package services

import (
	"errors"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
)

type CatalogService struct {
	CatalogRepo repository.CatalogRepository
	Auth        helpers.Auth
	Config      config.AppConfig
}

func (s *CatalogService) CreateCategory(category dto.CreateCategoryInput) error {
	err := s.CatalogRepo.CreateCategory(&models.Category{
		Name:         category.Name,
		ImageUrl:     category.ImageUrl,
		DisplayOrder: category.DisplayOrder,
	})
	return err
}

func (s *CatalogService) FindCategories() ([]*models.Category, error) {
	categories, err := s.CatalogRepo.FindCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *CatalogService) FindCategoryById(id uint) (*models.Category, error) {
	category, err := s.CatalogRepo.FindCategoryById(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CatalogService) EditCategory(id uint, category dto.CreateCategoryInput) (*models.Category, error) {
	_, err := s.CatalogRepo.FindCategoryById(id)
	if err != nil {
		return nil, errors.New("failed to find category")
	}

	updatedCategory, err := s.CatalogRepo.EditCategory(&models.Category{
		Name:         category.Name,
		ImageUrl:     category.ImageUrl,
		DisplayOrder: category.DisplayOrder,
	})
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

func (s *CatalogService) DeleteCategory(id uint) error {
	category, err := s.CatalogRepo.FindCategoryById(id)
	if err != nil {
		return errors.New("failed to find category to delete")
	}

	err = s.CatalogRepo.DeleteCategory(category)
	if err != nil {
		return err
	}
	return nil
}
