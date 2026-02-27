package seeders

import (
	"gawean-be-go/internal/config"
	"gawean-be-go/internal/models"
)

func SeedLevel() {
	levels := []models.LevelModel{
		{LevelName: "Administrator", LevelCode: "ADM"},
		{LevelName: "Supervisor", LevelCode: "SPV"},
		{LevelName: "User", LevelCode: "USR"},
	}

	for _, level := range levels {
		config.DB.FirstOrCreate(&level, models.LevelModel{LevelCode: level.LevelCode})
	}
}
