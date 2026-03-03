package service

import (
	"errors"
	"strings"
	"time"

	"gawean-be-go/internal/models"
	"gawean-be-go/internal/repository"
	"gawean-be-go/pkg/utils"
)

type AuthService interface {
	Register(name, email, password string) error
	Login(email, password string) (string, string, *models.UserModel, error)
	Refresh(refreshToken string) (string, error)
	Logout(refreshToken string) error
}

type authService struct {
	userRepo    repository.UserRepository
	refreshRepo repository.RefreshTokenRepository
}

func NewAuthService(userRepo repository.UserRepository, refreshRepo repository.RefreshTokenRepository) AuthService {
	return &authService{userRepo, refreshRepo}
}

func (s *authService) Register(name, email, password string) error {
	email = strings.ToLower(strings.TrimSpace(email))

	existing, _ := s.userRepo.FindByEmail(email)
	if existing != nil && existing.UserID != 0 {
		return errors.New("email already registered")
	}

	user := &models.UserModel{
		Name:     name,
		Email:    email,
		Password: utils.HashPassword(password),
		LevelID:  1, // default level
	}

	return s.userRepo.Create(user)
}

func (s *authService) Login(email, password string) (string, string, *models.UserModel, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil || !utils.CheckPassword(user.Password, password) {
		return "", "", nil, errors.New("invalid credentials")
	}

	accessToken, _ := utils.GenerateAccessToken(user.UserID, user.Level.LevelCode)
	refreshToken, _ := utils.GenerateRefreshToken(user.UserID)

	refresh := &models.RefreshToken{
		UserID:    user.UserID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	s.refreshRepo.Create(refresh)

	return accessToken, refreshToken, user, nil
}

func (s *authService) Refresh(refreshToken string) (string, error) {

	stored, err := s.refreshRepo.FindValid(refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	if stored.ExpiresAt.Before(time.Now()) {
    return "", errors.New("refresh token expired")
}

	token, claims, err := utils.ParseToken(refreshToken)
	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	userID := uint(claims["user_id"].(float64))

	user, _ := s.userRepo.FindByID(userID)

	newAccess, _ := utils.GenerateAccessToken(userID, user.Level.LevelCode)

	return newAccess, nil
}

func (s *authService) Logout(refreshToken string) error {
	return s.refreshRepo.Revoke(refreshToken)
}