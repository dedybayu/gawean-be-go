package seeders

import (
	"gawean-be-go/internal/config"
	"gawean-be-go/internal/models"
)

func SeedLevel() {
	levels := []models.Level{
		{NamaLevel: "Administrator", Kode: "ADM"},
		{NamaLevel: "User", Kode: "USR"},
	}

	for _, level := range levels {
		config.DB.FirstOrCreate(&level, models.Level{Kode: level.Kode})
	}
}
