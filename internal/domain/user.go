package domain

import "time"

type User struct {
	UserID         uint      `json:"user_id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	LevelCode      string    `json:"level_code"`
	LevelName      string    `json:"level_name"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
