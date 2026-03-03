package service

import (
	"gawean-be-go/internal/domain"
	"gawean-be-go/internal/models"
	"gawean-be-go/internal/repository"
)

type UserService interface {
	GetUserInfo(userID uint) (*models.UserModel, error)
	GetAllUsers() ([]domain.User, error)
	GetUserProfile(userID uint) (*domain.UserProfile, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) GetUserInfo(userID uint) (*models.UserModel, error) {
	return s.userRepo.FindByID(userID)
}

func (s *userService) GetAllUsers() ([]domain.User, error) {
	return s.userRepo.FindAll()
}

// Get User Profile
func (s *userService) GetUserProfile(userID uint) (*domain.UserProfile, error) {
	return s.userRepo.GetUserProfile(userID)
}