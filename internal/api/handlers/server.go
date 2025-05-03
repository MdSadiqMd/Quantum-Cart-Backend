package handlers

import (
	"log"
	"net/http"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/api/utils"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.Open(postgres.Open(config.DataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("error in db connection: %v", err)
	}
	log.Println("Database connected successfully 🚀")

	db.AutoMigrate(&models.User{})
	auth := helpers.NewAuth(config.AppSecret)

	app.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"status": "OK",
		})
	})

	apiHandler := &utils.Handler{
		App:  app,
		DB:   db,
		Auth: auth,
	}
	SetupRoutes(apiHandler)

	log.Fatal(app.Listen(config.ServerPort))
}

func SetupRoutes(handler *utils.Handler) {
	SetupUserRoutes(handler)
}
