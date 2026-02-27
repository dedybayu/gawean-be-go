package repository

import (
	"gawean-be-go/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id uint) (*models.UserModel, error)
	FindAll() ([]models.UserModel, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByID(id uint) (*models.UserModel, error) {
	var user models.UserModel
	err := r.db.Preload("Level").First(&user, id).Error
	
	return &user, err
}

func (r *userRepository) FindAll() ([]models.UserModel, error) {
	var users []models.UserModel
	err := r.db.Preload("Level").Find(&users).Error
	return users, err
}