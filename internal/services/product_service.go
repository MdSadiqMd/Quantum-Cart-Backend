package services

import (
	"time"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
	"github.com/gofiber/fiber/v2"
)

type ProductService struct {
	ProductRepo repository.ProductRepository
	Auth        helpers.Auth
	Config      config.AppConfig
}

func NewProductService(productRepo repository.ProductRepository, auth helpers.Auth, config config.AppConfig) *ProductService {
	return &ProductService{
		ProductRepo: productRepo,
		Auth:        auth,
		Config:      config,
	}
}

func (s *ProductService) CreateProduct(ctx *fiber.Ctx, input dto.CreateProductRequest) (*models.Product, error) {
	user := s.Auth.GetCurrentUser(ctx)
	product, err := s.ProductRepo.CreateProduct(&models.Product{
		Name:        input.Name,
		Description: input.Description,
		CategoryId:  input.CategoryId,
		ImageUrls:   input.ImageUrls,
		Price:       input.Price,
		UserId:      int(user.Id),
		Stock:       input.Stock,
	})
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) FindProducts(limit, offset int) ([]*models.Product, error) {
	products, err := s.ProductRepo.FindProducts(limit, offset)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) FindProductbyId(id uint) (*models.Product, error) {
	product, err := s.ProductRepo.FindProductbyId(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) FindSellerProducts(id uint) ([]*models.Product, error) {
	products, err := s.ProductRepo.FindSellerProducts(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) UpdateProduct(id uint, input dto.CreateProductRequest, user *models.User) (*models.Product, error) {
	existingProduct, err := s.ProductRepo.FindProductbyId(id)
	if err != nil {
		return nil, err
	}

	existingProduct.Name = input.Name
	existingProduct.Description = input.Description
	existingProduct.CategoryId = input.CategoryId
	existingProduct.ImageUrls = input.ImageUrls
	existingProduct.Price = input.Price
	existingProduct.Stock = input.Stock
	existingProduct.UserId = int(user.Id)
	existingProduct.UpdatedAt = time.Now()

	updatedProduct, err := s.ProductRepo.UpdateProduct(id, existingProduct)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (s *ProductService) UpdateProductStock(id uint, stock uint) (*models.Product, error) {
	existingProduct, err := s.ProductRepo.FindProductbyId(id)
	if err != nil {
		return nil, err
	}

	existingProduct.Stock = stock
	updatedProduct, err := s.ProductRepo.UpdateProduct(id, existingProduct)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (s *ProductService) DeleteProduct(id uint) (*models.Product, error) {
	product, err := s.ProductRepo.FindProductbyId(id)
	if err != nil {
		return nil, err
	}

	deleteProduct, err := s.ProductRepo.DeleteProduct(product)
	if err != nil {
		return nil, err
	}
	return deleteProduct, nil
}
