package domain

import "time"

type Level struct {
	LevelID   uint
	LevelName string
	LevelCode string
	CreatedAt time.Time
	UpdatedAt time.Time
}
