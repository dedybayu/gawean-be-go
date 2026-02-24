package config

import "log"

func ResetDatabase() {
	err := DB.Exec(`
		DROP TABLE IF EXISTS users CASCADE;
		DROP TABLE IF EXISTS levels CASCADE;
	`).Error

	if err != nil {
		log.Fatal("Failed to reset database:", err)
	}

	log.Println("Database reset completed")
}
