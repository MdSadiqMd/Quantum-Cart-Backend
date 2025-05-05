package handlers

import (
	"net/http"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/api/utils"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/services"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/response"
	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	service services.CatalogService
}

func SetupCatalogRoutes(handler *utils.Handler) {
	app := handler.App

	catalogService := services.CatalogService{
		CatalogRepo: repository.NewCatalogRepository(handler.DB),
		Auth:        handler.Auth,
		Config:      handler.Config,
	}

	catalogHandler := CatalogHandler{
		service: catalogService,
	}

	publicRoutes := app.Group("/")
	publicRoutes.Get("/products")
	publicRoutes.Get("/products/:id")
	publicRoutes.Get("/categories")
	publicRoutes.Get("/categories/:id")

	privateRoutes := publicRoutes.Group("/seller", handler.Auth.AuthorizeSeller)
	privateRoutes.Post("/categories", catalogHandler.CreateCategory)
	privateRoutes.Put("/categories/:id", catalogHandler.EditCategory)
	privateRoutes.Delete("/categories/:id", catalogHandler.DeleteCategory)
	privateRoutes.Post("/products", catalogHandler.CreateProduct)
	privateRoutes.Put("/products/:id", catalogHandler.EditProduct)
	privateRoutes.Patch("/products/:id", catalogHandler.UpdateStock)
	privateRoutes.Delete("/products/:id", catalogHandler.DeleteProduct)
}

func (h CatalogHandler) CreateCategory(ctx *fiber.Ctx) error {
	return response.SuccessResponse(ctx, http.StatusOK, "Category created successfully", nil)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {
	return response.SuccessResponse(ctx, http.StatusOK, "Category edited successfully", nil)
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	return response.SuccessResponse(ctx, http.StatusOK, "Category deleted successfully", nil)
}

func (h CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {
	return response.SuccessResponse(ctx, http.StatusOK, "Product created successfully", nil)
}

func (h CatalogHandler) EditProduct(ctx *fiber.Ctx) error {
	return response.SuccessResponse(ctx, http.StatusOK, "Product edited successfully", nil)
}

func (h CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {
	return response.SuccessResponse(ctx, http.StatusOK, "Product stock updated successfully", nil)
}

func (h CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {
	return response.SuccessResponse(ctx, http.StatusOK, "Product deleted successfully", nil)
}
