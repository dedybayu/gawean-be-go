package repository

import (
	"gawean-be-go/internal/models"

	"gorm.io/gorm"
)

type LevelRepository interface {
	FindByCode(code string) (*models.LevelModel, error)
}

type levelRepository struct {
	db *gorm.DB
}

func NewLevelRepository(db *gorm.DB) LevelRepository {
	return &levelRepository{db}
}

func (r *levelRepository) FindByCode(code string) (*models.LevelModel, error) {
	var level models.LevelModel
	err := r.db.Where("level_code = ?", code).First(&level).Error
	if err != nil {
		return nil, err
	}
	return &level, nil
}