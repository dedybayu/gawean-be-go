package models

import "time"

type RefreshToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint
	Token     string    `gorm:"uniqueIndex"`
	ExpiresAt time.Time
	Revoked   bool      `gorm:"default:false"`
	CreatedAt time.Time
}