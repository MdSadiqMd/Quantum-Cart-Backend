package repository

import (
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"gorm.io/gorm"
)

type BankRepository interface {
	CreateBankAccount(bank models.BankAccount) error
}

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {
	return &bankRepository{
		db: db,
	}
}

func (r bankRepository) CreateBankAccount(bank models.BankAccount) error {
	err := r.db.Create(&bank).Error
	if err != nil {
		return err
	}
	return nil
}
