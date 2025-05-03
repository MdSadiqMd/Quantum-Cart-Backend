package services

import (
	"errors"
	"log"
	"time"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
	Auth     helpers.Auth
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {
	hashedPassword, err := s.Auth.CreateHashedPassowrd(input.Password)
	if err != nil {
		return "", errors.New("failed to hash password")
	}

	user, err := s.UserRepo.CreateUser(models.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	})
	if err != nil {
		log.Printf("error in creating user: %v", err)
		return "", errors.New("failed to create user")
	}

	return s.Auth.GenerateToken(user.Id, user.Email, user.UserType)
}

func (s UserService) Login(email string, password string) (string, error) {
	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("failed to find user")
	}

	err = s.Auth.VerifyPassword(password, user.Password)
	if err != nil {
		return "", errors.New("password does not match")
	}

	return s.Auth.GenerateToken(user.Id, user.Email, user.UserType)
}

func (s UserService) findUserByEmail(email string) (*models.User, error) {
	user, err := s.UserRepo.FindUser(email)
	return &user, err
}

func (s UserService) isVerifiedUser(id uint) bool {
	user, err := s.UserRepo.FindUserById(id)
	return err != nil && user.Verified
}

func (s UserService) GetVerificationCode(user models.User) (int, error) {
	if s.isVerifiedUser(user.Id) {
		return user.Code, nil
	}

	code, err := s.Auth.GenerateCode()
	if err != nil {
		return 0, errors.New("failed to generate code")
	}

	updateUser := models.User{
		Expiry: time.Now().Add(time.Minute * 30),
		Code:   code,
	}

	_, err = s.UserRepo.UpdateUser(user.Id, updateUser)
	if err != nil {
		return 0, errors.New("failed to update verification code")
	}
	return code, nil
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
