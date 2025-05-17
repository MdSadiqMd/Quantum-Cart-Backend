package handlers

import (
	"log"
	"net/http"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/api/utils"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/payment"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	err = db.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.BankAccount{},
		&models.Category{},
		&models.Cart{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatalf("error in db migration: %v", err)
	}
	log.Println("Database migrated successfully 🔀")

	corsConfig := cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Accept, Authorization, Content-Type",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	})
	app.Use(corsConfig)

	auth := helpers.NewAuth(config.AppSecret)
	paymentClient := payment.NewPaymentClient(config.StripeSecret, config.SuccessURL, config.CancelURL)
	app.Get("/healthz", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"status": "OK",
		})
	})

	apiHandler := &utils.Handler{
		App:           app,
		DB:            db,
		Auth:          auth,
		Config:        config,
		PaymentClient: paymentClient,
	}
	SetupRoutes(apiHandler)

	log.Fatal(app.Listen(config.ServerPort))
}

func SetupRoutes(handler *utils.Handler) {
	SetupUserRoutes(handler)
	SetupCatalogRoutes(handler)
	SetupProductRoutes(handler)
	SetupCartRoutes(handler)
	SetupOrderRoutes(handler)
}
