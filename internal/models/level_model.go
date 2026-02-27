package models

import "time"

type LevelModel struct {
	LevelID   uint   `gorm:"primaryKey" json:"level_id"`
	LevelName string `gorm:"size:50;not null" json:"level_name"`
	LevelCode string `gorm:"size:10;uniqueIndex;not null" json:"level_code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (LevelModel) TableName() string {
	return "levels"
}