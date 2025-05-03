package handlers

import (
	"net/http"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/api/utils"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service services.UserService
	auth    helpers.Auth
}

func SetupUserRoutes(handler *utils.Handler) {
	app := handler.App

	userService := services.UserService{
		UserRepo: repository.NewUserRepository(handler.DB),
		Auth:     handler.Auth,
	}
	userHandler := UserHandler{
		service: userService,
	}

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)

	app.Get("/verify", userHandler.GetVerificationCode)
	app.Post("/verify", userHandler.Verify)
	app.Post("/profile", userHandler.CreateProfile)
	app.Get("/profile", userHandler.GetProfile)
	app.Post("/cart", userHandler.AddToCart)
	app.Get("/cart", userHandler.GetCart)
	app.Post("/order", userHandler.CreateOrder)
	app.Get("/orders", userHandler.GetOrders)
	app.Get("/order/:id", userHandler.GetOrder)
	app.Post("/become-seller", userHandler.BecomeSeller)
}

func (h *UserHandler) Register(ctx *fiber.Ctx) error {
	user := dto.UserSignup{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "Failed at Backend DTO",
		})
	}

	token, err := h.service.Signup(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": "Error at Signup Auth",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": token,
	})
}

func (h *UserHandler) Login(ctx *fiber.Ctx) error {
	user := dto.UserLogin{}
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": "Failed at Backend DTO",
		})
	}

	token, err := h.service.Login(user.Email, user.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"error": "Error at Login Auth",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success": token,
	})
}

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User get verification code successfully",
	})
}

func (h *UserHandler) Verify(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User verified successfully",
	})
}

func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User profile created successfully",
	})
}

func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User profile fetched successfully",
	})
}

func (h *UserHandler) AddToCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User added to cart successfully",
	})
}

func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User cart fetched successfully",
	})
}

func (h *UserHandler) CreateOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User order created successfully",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User orders fetched successfully",
	})
}

func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User order fetched successfully",
	})
}

func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User become seller successfully",
	})
}
