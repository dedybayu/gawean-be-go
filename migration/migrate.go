package migration

import (
	"gawean-be-go/internal/models"
	"gawean-be-go/internal/config"
	"log"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&models.LevelModel{},
		&models.UserModel{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration success")
}
