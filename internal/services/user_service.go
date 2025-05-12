package services

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/dto"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/helpers"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/repository"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/events"
)

type UserService struct {
	UserRepo repository.UserRepository
	BankRepo repository.BankRepository
	Auth     helpers.Auth
	Config   config.AppConfig
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
	if err != nil {
		return nil, errors.New("failed to find user")
	}
	return &user, nil
}

func (s UserService) isVerifiedUser(id uint) bool {
	user, err := s.UserRepo.FindUserById(id)
	return err != nil && user.Verified
}

func (s UserService) GetVerificationCode(user models.User) error {
	if s.isVerifiedUser(user.Id) {
		return nil
	}

	code, err := s.Auth.GenerateCode()
	if err != nil {
		return errors.New("failed to generate code")
	}

	updateUser := models.User{
		Expiry: time.Now().Add(time.Minute * 30),
		Code:   code,
	}

	_, err = s.UserRepo.UpdateUser(user.Id, updateUser)
	if err != nil {
		return errors.New("failed to update verification code")
	}

	user, err = s.UserRepo.FindUserById(user.Id)
	if err != nil {
		return errors.New("failed to find user")
	}

	notificationClient := events.NewNotificationClient(s.Config)
	err = notificationClient.SendSMS(user.Phone, "Your verification code is: "+strconv.Itoa(code))
	if err != nil {
		return errors.New("failed to send verification code")
	}
	return nil
}

func (s UserService) VerifyCode(id uint, code int) error {
	if s.isVerifiedUser(id) {
		return errors.New("user is already verified")
	}

	user, err := s.UserRepo.FindUserById(id)
	if err != nil {
		return errors.New("failed to find user")
	}
	if user.Code != code {
		return errors.New("invalid verification code")
	}
	if time.Now().After(user.Expiry) {
		return errors.New("verification code is expired")
	}

	updateUser := models.User{
		Verified: true,
	}

	_, err = s.UserRepo.UpdateUser(user.Id, updateUser)
	if err != nil {
		return errors.New("failed to update verification code")
	}
	return nil
}

func (s UserService) CreateProfile(id uint, user dto.ProfileInput) (models.Address, error) {
	_, err := s.UserRepo.UpdateUser(id, models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		return models.Address{}, errors.New("failed to update user")
	}

	address := models.Address{
		AddressLine1: user.AddressInput.AddressLine1,
		AddressLine2: user.AddressInput.AddressLine2,
		City:         user.AddressInput.City,
		Country:      user.AddressInput.Country,
		PostCode:     user.AddressInput.PostCode,
		UserId:       id,
	}

	address, err = s.UserRepo.CreateProfile(address)
	if err != nil {
		return models.Address{}, errors.New("failed to create profile")
	}
	return address, nil
}

func (s UserService) GetProfile(id uint) (*models.User, error) {
	user, err := s.UserRepo.FindUserById(id)
	if err != nil {
		return nil, errors.New("failed to find user")
	}
	return &user, nil
}

func (s UserService) UpdateProfile(id uint, user dto.ProfileInput) (models.Address, error) {
	existingUser, err := s.UserRepo.FindUserById(id)
	if err != nil {
		return models.Address{}, errors.New("failed to find user")
	}

	if existingUser.FirstName != user.FirstName {
		_, err := s.UserRepo.UpdateUser(id, models.User{
			FirstName: user.FirstName,
		})
		if err != nil {
			return models.Address{}, errors.New("failed to update user")
		}
	}

	if existingUser.LastName != user.LastName {
		_, err := s.UserRepo.UpdateUser(id, models.User{
			LastName: user.LastName,
		})
		if err != nil {
			return models.Address{}, errors.New("failed to update user")
		}
	}

	address, err := s.UserRepo.UpdateProfile(models.Address{
		AddressLine1: user.AddressInput.AddressLine1,
		AddressLine2: user.AddressInput.AddressLine2,
		City:         user.AddressInput.City,
		Country:      user.AddressInput.Country,
		PostCode:     user.AddressInput.PostCode,
		UserId:       id,
	})
	if err != nil {
		return models.Address{}, errors.New("failed to update profile")
	}
	return address, nil
}

func (s UserService) BecomeSeller(id uint, input dto.SellerInput) (string, error) {
	user, err := s.UserRepo.FindUserById(id)
	if err != nil {
		return "", errors.New("failed to find user")
	}
	if user.UserType == models.SELLER {
		return "", errors.New("user is already a seller")
	}

	seller, err := s.UserRepo.UpdateUser(id, models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.Phone,
		UserType:  models.SELLER,
	})
	if err != nil {
		return "", errors.New("failed to update user")
	}

	token, err := s.Auth.GenerateToken(user.Id, user.Email, seller.UserType)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	err = s.BankRepo.CreateBankAccount(models.BankAccount{
		BankAccount: input.BankAccountNumber,
		SwiftCode:   input.SwiftCode,
		PaymentType: input.PaymentType,
		UserId:      id,
	})
	if err != nil {
		log.Printf("Failed to create bank account: %v", err)
		return token, nil
	}

	return token, nil
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
