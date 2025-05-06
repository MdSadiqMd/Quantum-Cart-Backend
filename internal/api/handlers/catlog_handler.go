package handlers

import (
	"net/http"
	"strconv"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/api/utils"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
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
	publicRoutes.Get("/products", catalogHandler.GetProducts)
	publicRoutes.Get("/products/:id", catalogHandler.GetProduct)
	publicRoutes.Get("/categories", catalogHandler.GetCategories)
	publicRoutes.Get("/categories/:id", catalogHandler.GetCategory)

	privateRoutes := publicRoutes.Group("/seller", handler.Auth.AuthorizeSeller)
	privateRoutes.Post("/categories", catalogHandler.CreateCategory)
	privateRoutes.Patch("/categories/:id", catalogHandler.EditCategory)
	privateRoutes.Delete("/categories/:id", catalogHandler.DeleteCategory)
	privateRoutes.Post("/products", catalogHandler.CreateProduct)
	privateRoutes.Put("/products/:id", catalogHandler.EditProduct)
	privateRoutes.Patch("/products/:id", catalogHandler.UpdateStock)
	privateRoutes.Delete("/products/:id", catalogHandler.DeleteProduct)
}

func (h CatalogHandler) CreateCategory(ctx *fiber.Ctx) error {
	req := dto.CreateCategoryInput{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Category Input Data is not valid", err)
	}

	err = h.service.CreateCategory(req)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create category", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Category created successfully", req)
}

func (h CatalogHandler) GetCategories(ctx *fiber.Ctx) error {
	categories, err := h.service.FindCategories()
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch categories", err)
	}

	return response.SuccessResponse(ctx, http.StatusOK, "Categories fetched successfully", categories)
}

func (h CatalogHandler) GetCategory(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid category ID", err)
	}

	category, err := h.service.FindCategoryById(uint(id))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch category", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Category fetched successfully", category)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid category ID", err)
	}

	req := dto.CreateCategoryInput{}
	err = ctx.BodyParser(&req)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Category Input Data is not valid", err)
	}

	category, err := h.service.EditCategory(uint(id), req)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to edit category", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Category edited successfully", category)
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid category ID", err)
	}

	err = h.service.DeleteCategory(uint(id))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete category", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Category deleted successfully", nil)
}

func (h CatalogHandler) CreateProduct(ctx *fiber.Ctx) error {
	return response.SuccessResponse(ctx, http.StatusOK, "Product created successfully", nil)
}

func (h CatalogHandler) GetProducts(ctx *fiber.Ctx) error {
	return response.SuccessResponse(ctx, http.StatusOK, "Products fetched successfully", nil)
}

func (h CatalogHandler) GetProduct(ctx *fiber.Ctx) error {
	return response.SuccessResponse(ctx, http.StatusOK, "Product fetched successfully", nil)
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
