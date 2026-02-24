package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Nama      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:100;uniqueIndex;not null"`
	Password  string    `json:"-"`
	LevelID   uint      `gorm:"not null"`
	Level     Level     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
