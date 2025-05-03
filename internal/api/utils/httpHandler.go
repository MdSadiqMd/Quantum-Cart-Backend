package utils

import (
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	App    *fiber.App
	DB     *gorm.DB
	Auth   helpers.Auth
	Config config.AppConfig
}
