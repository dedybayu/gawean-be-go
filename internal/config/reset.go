package config

import "log"

func ResetDatabase() {
	err := DB.Exec(`
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
	`).Error

	if err != nil {
		log.Fatal("Failed to reset database:", err)
	}

	log.Println("All tables dropped successfully")
}