package models

import "time"

type User struct {
	UserID         uint
	Name           string
	Email          string
	Password       string
	LevelID        uint
	Level          Level
	ProfilePicture string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
