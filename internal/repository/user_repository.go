package repository

import (
	"gawean-be-go/internal/domain"
	"gawean-be-go/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id uint) (*models.UserModel, error)
	FindAll() ([]domain.User, error)
	FindByEmail(email string) (*models.UserModel, error)
	Create(user *models.UserModel) error
	Update(user *models.UserModel) error
	Delete(id uint) error
	GetUserProfile(id uint) (*domain.UserProfile, error)
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

func (r *userRepository) FindAll() ([]domain.User, error) {
	var users []models.UserModel
	err := r.db.Preload("Level").Find(&users).Error

	var result []domain.User
	for _, u := range users {
		result = append(result, domain.User{
			UserID:         u.UserID,
			Name:           u.Name,
			Email:          u.Email,
			LevelCode:      u.Level.LevelCode,
			LevelName:      u.Level.LevelName,
			ProfilePicture: u.ProfilePicture,
			CreatedAt:      u.CreatedAt,
			UpdatedAt:      u.UpdatedAt,
		})
	}
	return result, err
}

func (r *userRepository) FindByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel
	err := r.db.Preload("Level").
		Where("email = ?", email).
		First(&user).Error
	return &user, err
}

func (r *userRepository) Create(user *models.UserModel) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *models.UserModel) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.UserModel{}, id).Error
}

// Get Profile
func (r *userRepository) GetUserProfile(id uint) (*domain.UserProfile, error) {
	var user models.UserModel
	err := r.db.Preload("Level").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &domain.UserProfile{
		UserID:         user.UserID,
		Name:           user.Name,
		Email:          user.Email,
		LevelCode:      user.Level.LevelCode,
		LevelName:      user.Level.LevelName,
		ProfilePicture: user.ProfilePicture,
	}, nil
}
