package service

import (
	"gawean-be-go/internal/models"
	"gawean-be-go/internal/repository"
)

type UserService interface {
	GetUserInfo(userID uint) (*models.UserModel, error)
	GetAllUsers() ([]models.UserModel, error)
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

func (s *userService) GetAllUsers() ([]models.UserModel, error) {
	return s.userRepo.FindAll()
}