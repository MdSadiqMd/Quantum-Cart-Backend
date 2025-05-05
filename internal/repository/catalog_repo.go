package repository

import "gorm.io/gorm"

type CatalogRepository interface {
	CreateCatalog(catalog models.Catalog) error
	FindCatalog(id uint) (models.Catalog, error)
	FindAllCatalogs() ([]models.Catalog, error)
	UpdateCatalog(id uint, catalog models.Catalog) (models.Catalog, error)
	DeleteCatalog(catalog models.Catalog) (models.Catalog, error)
}

type catalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{
		db: db,
	}
}


