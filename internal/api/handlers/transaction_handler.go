package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/api/utils"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/services"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/payment"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/response"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	service       services.TransactionService
	orderService  services.OrderService
	cartService   services.CartService
	paymentClient payment.PaymentClient
}

func SetupTransactionRoutes(handler *utils.Handler) {
	app := handler.App

	transactionService := services.TransactionService{
		TransactionRepo: repository.NewTransactionRepository(handler.DB),
		UserRepo:        repository.NewUserRepository(handler.DB),
		OrderRepo:       repository.NewOrderRepository(handler.DB),
		Auth:            handler.Auth,
		Config:          handler.Config,
	}

	orderService := services.OrderService{
		OrderRepo: repository.NewOrderRepository(handler.DB),
		Config:    handler.Config,
	}

	cartService := services.CartService{
		CartRepo: repository.NewCartRepository(handler.DB),
	}

	transactionHandler := TransactionHandler{
		service:       transactionService,
		orderService:  orderService,
		cartService:   cartService,
		paymentClient: handler.PaymentClient,
	}

	publicRoutes := app.Group("/")
	privateRoutes := publicRoutes.Group("/transaction", handler.Auth.Authorize)
	privateRoutes.Get("/payment", transactionHandler.MakePayment)
	privateRoutes.Get("/verify", transactionHandler.VerifyPayment)
}

func (h *TransactionHandler) MakePayment(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)
	publishableKey := h.service.Config.PublishableKey

	_, amount, err := h.orderService.GetCurrentOrder(user.Id)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get active order", err)
	}

	orderId, err := helpers.RandomNumbers(8)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to generate order ID", err)
	}

	paymentResult, err := h.paymentClient.CreatePayment(amount, user.Id, uint(orderId))
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create payment", err)
	}

	createdTransaction, err := h.service.CreateTransaction(dto.CreatePaymentRequest{
		UserId:       user.Id,
		Amount:       amount,
		Status:       models.PaymentPending,
		OrderId:      fmt.Sprintf("%d", orderId),
		PaymentId:    paymentResult.ID,
		ClientSecret: paymentResult.ClientSecret,
	})
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create transaction", err)
	}

	return response.SuccessResponse(ctx, http.StatusOK, "Payment intent created successfully", fiber.Map{
		"publishableKey": publishableKey,
		"activePayment":  createdTransaction,
	})
}

func (h *TransactionHandler) VerifyPayment(ctx *fiber.Ctx) error {
	user := h.service.Auth.GetCurrentUser(ctx)

	activatePayment, err := h.service.GetActivePayment(user.Id)
	if err != nil || activatePayment.ID == 0 {
		return response.ErrorResponse(ctx, http.StatusNotFound, "No active payment found", err)
	}

	paymentResult, err := h.paymentClient.GetPaymentStatus(activatePayment.PaymentId)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to verify payment", err)
	}

	paymentJson, _ := json.Marshal(paymentResult)
	paymentLogs := string(paymentJson)

	var paymentStatusValue models.PaymentStatus
	if paymentResult.Status == "succeeded" {
		paymentStatusValue = models.PaymentSuccess
	} else {
		paymentStatusValue = models.PaymentFailed
	}

	_, err = h.service.UpdatePaymentLog(user.Id, &paymentStatusValue, paymentLogs)
	if err != nil {
		return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to update payment log", err)
	}

	if paymentStatusValue == models.PaymentSuccess {
		cartItems, err := h.cartService.CartRepo.FindCartItems(user.Id)
		if err != nil {
			return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get cart items", err)
		}

		var orderRequests []*dto.CreateOrderRequest
		for _, item := range cartItems {
			orderRequests = append(orderRequests, &dto.CreateOrderRequest{
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
				Name:      item.Name,
				Price:     item.Price,
				ImageUrls: strings.Split(item.ImageUrl, ","),
				SellerId:  item.SellerId,
			})
		}

		_, _, err = h.orderService.CreateOrder(user.Id, orderRequests, true)
		if err != nil {
			return response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to create order", err)
		}
	}

	return response.SuccessResponse(ctx, http.StatusOK, "Payment verified successfully", &fiber.Map{
		"message": "Payment verified successfully",
		"status":  paymentStatusValue,
		"result":  paymentResult,
	})
}
