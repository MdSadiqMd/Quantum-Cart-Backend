package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	user, err := s.UserRepo.CreateUser(models.User{
		Email:    input.Email,
		Phone:    input.Phone,
		Password: input.Password,
	})
	if err != nil {
		log.Printf("error in creating user: %v", err)
		return "", errors.New("failed to create user")
	}

	userInfo := fmt.Sprintf("%v %v %v", user.Id, user.Email, user.UserType)
	return userInfo, nil
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("failed to find user")
	}
	if user.Password != password {
		return "", errors.New("invalid password")
	}
	return user.Email, nil
}

func (s UserService) findUserByEmail(email string) (*models.User, error) {
	user, err := s.UserRepo.FindUser(email)
	return &user, err
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
