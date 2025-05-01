package handlers

import (
	"log"
	"net/http"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/api/utils"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"status": "OK",
		})
	})

	apiHandler := &utils.Handler{App: app}
	SetupRoutes(apiHandler)

	log.Fatal(app.Listen(config.ServerPort))
}

func SetupRoutes(handler *utils.Handler) {
	SetupUserRoutes(handler)
}
