package services

import (
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
)

type TransactionService struct {
	TransactionRepo repository.TranscationRepository
	UserRepo        repository.UserRepository
	OrderRepo       repository.OrderRepository
	Auth            helpers.Auth
	Config          config.AppConfig
}

func NewTransactionService(
	transactionRepo repository.TranscationRepository,
	userRepo repository.UserRepository,
	orderRepo repository.OrderRepository,
	auth helpers.Auth,
	config config.AppConfig,
) *TransactionService {
	return &TransactionService{
		TransactionRepo: transactionRepo,
		UserRepo:        userRepo,
		OrderRepo:       orderRepo,
		Auth:            auth,
		Config:          config,
	}
}

func (s *TransactionService) CreateTransaction(transaction dto.CreatePaymentRequest) (*models.Payment, error) {
	payment := models.Payment{
		UserId:       transaction.UserId,
		Amount:       transaction.Amount,
		Status:       models.PaymentInit,
		OrderId:      transaction.OrderId,
		PaymentId:    transaction.PaymentId,
		ClientSecret: transaction.ClientSecret,
	}
	createPayment, err := s.TransactionRepo.CreatePayment(&payment)
	if err != nil {
		return nil, err
	}
	return createPayment, nil
}

func (s *TransactionService) GetPaymentStatus(userId uint) (*models.PaymentStatus, error) {
	payment, err := s.TransactionRepo.FindInitialPayment(userId)
	if err != nil {
		return nil, err
	}
	return &payment.Status, nil
}

func (s *TransactionService) GetActivePayment(userId uint) (*models.Payment, error) {
	payment, err := s.TransactionRepo.FindInitialPayment(userId)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *TransactionService) UpdatePaymentLog(userId uint, status *models.PaymentStatus, paymentLog string) (*models.Payment, error) {
	payment, err := s.TransactionRepo.FindInitialPayment(userId)
	if err != nil {
		return nil, err
	}
	payment.Status = *status
	payment.Response = paymentLog

	updatedPaymentLog, err := s.TransactionRepo.UpdatePayment(payment)
	if err != nil {
		return nil, err
	}
	return updatedPaymentLog, nil
}
