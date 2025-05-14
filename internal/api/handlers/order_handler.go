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

type OrderHandler struct {
	service services.OrderService
}

func SetupOrderRoutes(handler *utils.Handler) {
	app := handler.App

	orderService := services.OrderService{
		OrderRepo: repository.NewOrderRepository(handler.DB),
		CartRepo:  repository.NewCartRepository(handler.DB),
		Auth:      handler.Auth,
		Config:    handler.Config,
	}
	orderHandler := OrderHandler{
		service: orderService,
	}

	publicRoutes := app.Group("/")
	privateRoutes := publicRoutes.Group("/order", handler.Auth.Authorize)
	privateRoutes.Post("/", orderHandler.CreateOrder)
	privateRoutes.Get("/", orderHandler.GetOrders)
	privateRoutes.Get("/:id", orderHandler.GetOrder)
}

func (h *OrderHandler) CreateOrder(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)
	var req dto.BulkOrderRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "Order input data is not valid", err)
	}
	if len(req.Orders) == 0 {
		return response.ErrorResponse(ctx, http.StatusBadRequest, "No items provided in the order", nil)
	}

	orders, refs, err := h.service.CreateOrder(user.Id, req.Orders, req.IsBulk)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create order", err)
	}

	if req.IsBulk {
		return response.SuccessResponse(ctx, http.StatusOK, "Bulk order created successfully", dto.BulkOrderResponse{
			Orders:    orders,
			OrderRefs: refs,
		})
	}

	return response.SuccessResponse(ctx, http.StatusOK, "Order created successfully", dto.OrderResponse{
		Order:    orders[0],
		OrderRef: refs[0],
	})
}

func (h *OrderHandler) GetOrders(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)
	orders, err := h.service.GetOrders(user.Id)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch orders", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Orders fetched successfully", orders)
}

func (h *OrderHandler) GetOrder(ctx *fiber.Ctx) error {
	orderId, _ := strconv.Atoi(ctx.Params("id"))
	user := h.service.Auth.GetCurrentUser(ctx)
	order, err := h.service.GetOrderById(uint(orderId), user.Id)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to fetch order", err)
	}
	return response.SuccessResponse(ctx, http.StatusOK, "Order fetched successfully", order)
}
