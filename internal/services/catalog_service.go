package services

import (
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
)

type CatalogService struct {
	CatalogRepo repository.CatalogRepository
	Auth        helpers.Auth
	Config      config.AppConfig
}
