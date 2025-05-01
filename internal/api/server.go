package api

import (
	"log"
	"net/http"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
	"github.com/gofiber/fiber/v2"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	app.Get("/healthz", HealthCheck)

	log.Fatal(app.Listen(config.ServerPort))
}

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"status": "OK",
	})
}
