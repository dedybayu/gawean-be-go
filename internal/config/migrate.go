package config

import (
	"gawean-be-go/internal/models"
	"log"
)

func Migrate() {
	err := DB.AutoMigrate(
		&models.Level{},
		&models.User{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration success")
}
