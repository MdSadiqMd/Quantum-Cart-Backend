package services

import (
	"log"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/domain"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
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

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	return &domain.User{}, nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) VerifyCode(id uint, input any) error {
	return nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return &domain.User{}, nil
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

func (s UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]interface{}, error) {
	return nil, nil
}

func (s UserService) GetOrderById(id uint, userId uint) ([]interface{}, error) {
	return nil, nil
}
