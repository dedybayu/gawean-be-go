package repository

import (
	"gawean-be-go/internal/models"
	"time"

	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(token *models.RefreshToken) error
	FindValid(token string) (*models.RefreshToken, error)
	Revoke(token string) error
}

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepository{db}
}

func (r *refreshTokenRepository) Create(token *models.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *refreshTokenRepository) FindValid(token string) (*models.RefreshToken, error) {
	var stored models.RefreshToken
	err := r.db.
		Where("token = ? AND revoked = ?", token, false).
		First(&stored).Error

	if err != nil || time.Now().After(stored.ExpiresAt) {
		return nil, err
	}

	return &stored, nil
}

func (r *refreshTokenRepository) Revoke(token string) error {
	return r.db.Model(&models.RefreshToken{}).
		Where("token = ?", token).
		Update("revoked", true).Error
}