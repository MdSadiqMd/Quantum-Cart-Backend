package repository

import (
	"errors"
	"log"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	FindUser(email string) (models.User, error)
	FindUserById(id uint) (models.User, error)
	UpdateUser(id uint, user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
	CreateProfile(user models.Address) (models.Address, error)
	UpdateProfile(user models.Address) (models.Address, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		log.Printf("error in creating user: %v", err)
		return models.User{}, errors.New("failed to create user")
	}
	return user, nil
}

func (r *userRepository) FindUser(email string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("error in finding user: %v", err)
		return models.User{}, errors.New("failed to find user")
	}
	return user, nil
}

func (r *userRepository) FindUserById(id uint) (models.User, error) {
	var user models.User
	err := r.db.Preload("Address").
		Preload("Cart").
		Preload("Orders").
		Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Printf("error in finding user: %v", err)
		return models.User{}, errors.New("failed to find user")
	}
	return user, nil
}

func (r *userRepository) UpdateUser(id uint, user models.User) (models.User, error) {
	var updatedUser models.User
	result := r.db.Model(&models.User{}).Where("id = ?", id).Clauses(clause.Returning{}).Updates(user).Scan(&updatedUser)
	if result.Error != nil {
		log.Printf("error in updating user: %v", result.Error)
		return models.User{}, errors.New("failed to update user")
	}
	return updatedUser, nil
}

func (r *userRepository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error
	if err != nil {
		log.Printf("error in deleting user: %v", err)
		return models.User{}, errors.New("failed to delete user")
	}
	return user, nil
}

func (r *userRepository) CreateProfile(user models.Address) (models.Address, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		log.Printf("error in creating profile: %v", err)
		return models.Address{}, errors.New("failed to create profile")
	}
	return user, nil
}

func (r *userRepository) UpdateProfile(user models.Address) (models.Address, error) {
	err := r.db.Model(&models.Address{}).Where("user_id = ?", user.UserId).Clauses(clause.Returning{}).Updates(user).Error
	if err != nil {
		return models.Address{}, errors.New("failed to update profile")
	}
	return user, nil
}
