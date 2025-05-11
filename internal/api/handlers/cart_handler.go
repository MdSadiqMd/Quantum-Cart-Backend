package handlers

import (
	"net/http"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/api/utils"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/services"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/response"
	"github.com/gofiber/fiber/v2"
)

type CartHandler struct {
	service services.CartService
}

func SetupCartRoutes(handler *utils.Handler) {
	app := handler.App

	cartService := services.CartService{
		CartRepo:    repository.NewCartRepository(handler.DB),
		ProductRepo: repository.NewProductRepository(handler.DB),
		Auth:        handler.Auth,
		Config:      handler.Config,
	}

	cartHandler := CartHandler{
		service: cartService,
	}

	publicRoutes := app.Group("/")
	privateRoutes := publicRoutes.Group("/cart", handler.Auth.Authorize)
	privateRoutes.Post("/", cartHandler.AddToCart)
	privateRoutes.Get("/", cartHandler.GetCart)
}

func (h *CartHandler) AddToCart(ctx *fiber.Ctx) error {
	req := dto.CreateCartRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Cart Input Data is not valid", err)
	}

	user := h.service.Auth.GetCurrentUser(ctx)
	cartItems, err := h.service.CreateCart(req, user)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create cart", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Cart created successfully", cartItems)
}

func (h *CartHandler) GetCart(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)

	cart, _, err := h.service.FindCart(user.Id)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get cart", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Cart fetched successfully", cart)
}
