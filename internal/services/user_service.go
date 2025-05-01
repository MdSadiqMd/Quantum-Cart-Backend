package services

import (
	"log"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
)

type UserService struct {
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	log.Println(input)
	return "token", nil
}

func (s UserService) Login(input any) (string, error) {
	return "", nil
}

func (s UserService) findUserByEmail(email string) (*models.User, error) {
	return &models.User{}, nil
}

func (s UserService) GetVerificationCode(e models.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, input any) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*models.User, error) {
	return &models.User{}, nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}

func (s UserService) GetCart(id uint) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateCart(input any, u models.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(u models.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(u models.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, userId uint) ([]interface{}, error) {
	return nil, nil
}
