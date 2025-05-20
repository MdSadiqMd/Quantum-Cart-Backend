package repository

import (
	"errors"
	"log"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"gorm.io/gorm"
)

type TranscationRepository interface {
	CreatePayment(payment *models.Payment) (*models.Payment, error)
	FindInitialPayment(userId uint) (*models.Payment, error)
	UpdatePayment(payment *models.Payment) (*models.Payment, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TranscationRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) CreatePayment(payment *models.Payment) (*models.Payment, error) {
	err := r.db.Model(&models.Payment{}).Create(&payment).Error
	if err != nil {
		log.Println("error in creating payment: ", err)
		return nil, errors.New("failed to create payment")
	}
	return payment, nil
}

func (r *transactionRepository) FindInitialPayment(userId uint) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.Where("user_id = ? AND status = ?", userId, "initial").First(&payment).Order("created_at DESC").Error
	if err != nil {
		log.Println("error in finding initial payment: ", err)
		return nil, errors.New("failed to find initial payment")
	}
	return &payment, nil
}

func (r *transactionRepository) UpdatePayment(payment *models.Payment) (*models.Payment, error) {
	err := r.db.Model(&models.Payment{}).Where("id = ?", payment.ID).Updates(&payment).Error
	if err != nil {
		log.Println("error in updating payment: ", err)
		return nil, errors.New("failed to update payment")
	}
	return payment, nil
}
