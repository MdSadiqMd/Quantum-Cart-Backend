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

type ProductHandler struct {
	service services.ProductService
}

func SetupProductRoutes(handler *utils.Handler) {
	app := handler.App

	productService := services.ProductService{
		ProductRepo: repository.NewProductRepository(handler.DB),
		Auth:        handler.Auth,
		Config:      handler.Config,
	}

	productHandler := ProductHandler{
		service: productService,
	}

	publicRoutes := app.Group("/")
	publicRoutes.Get("/products", productHandler.GetProducts)
	publicRoutes.Get("/product/:id", productHandler.GetProductById)

	privateRoutes := publicRoutes.Group("/seller", handler.Auth.AuthorizeSeller)
	privateRoutes.Post("/product", productHandler.CreateProduct)
	privateRoutes.Get("/:id", productHandler.GetProductSeller)
	privateRoutes.Put("/product/:id", productHandler.EditProduct)
	privateRoutes.Patch("/product/:id", productHandler.UpdateProductStock)
	privateRoutes.Delete("/product/:id", productHandler.DeleteProduct)
}

func (h ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	req := dto.CreateProductRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Product Input Data is not valid", err)
	}

	product, err := h.service.CreateProduct(ctx, req)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create product", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Product created successfully", product)
}

func (h ProductHandler) GetProducts(ctx *fiber.Ctx) error {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid limit or offset", err)
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid limit or offset", err)
	}

	products, err := h.service.FindProducts(limit, offset)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch products", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Products fetched successfully", products)
}

func (h ProductHandler) GetProductById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid product ID", err)
	}

	product, err := h.service.FindProductbyId(uint(id))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch product", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Product fetched successfully", product)
}

func (h ProductHandler) GetProductSeller(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid product ID", err)
	}

	product, err := h.service.FindSellerProducts(uint(id))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch product", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Product fetched successfully", product)
}

func (h ProductHandler) EditProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid product ID", err)
	}

	req := dto.CreateProductRequest{}
	err = ctx.BodyParser(&req)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Product Input Data is not valid", err)
	}

	user := h.service.Auth.GetCurrentUser(ctx)
	product, err := h.service.UpdateProduct(uint(id), req, &user)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to edit product", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Product edited successfully", product)
}

func (h ProductHandler) UpdateProductStock(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid product ID", err)
	}

	req := dto.UpdateStockRequest{}
	err = ctx.BodyParser(&req)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Product Input Data is not valid", err)
	}

	product, err := h.service.UpdateProductStock(uint(id), req.Stock)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update product", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Product Stock updated successfully", product)
}

func (h ProductHandler) DeleteProduct(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid product ID", err)
	}

	_, err = h.service.FindProductbyId(uint(id))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch product", err)
	}

	product, err := h.service.DeleteProduct(uint(id))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to delete product", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Product deleted successfully", product)
}
