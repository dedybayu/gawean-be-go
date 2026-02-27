package models

import "time"

type UserModel struct {
	UserID         uint       `gorm:"primaryKey" json:"user_id"`
	Name           string     `gorm:"size:100;not null" json:"name"`
	Email          string     `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password       string     `json:"-"`
	LevelID        uint       `gorm:"not null" json:"level_id"`
	Level          LevelModel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"level"`
	ProfilePicture string     `gorm:"size:255" json:"profile_picture"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (UserModel) TableName() string {
	return "users"
}
