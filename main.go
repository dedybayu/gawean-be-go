package main

import (
	"flag"
	"log"
	"os"

	"gawean-be-go/internal/config"
	"gawean-be-go/internal/routes"
	"gawean-be-go/internal/seeders"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// ===== CLI FLAGS =====
	migrate := flag.Bool("migrate", false, "Run database migration")
	seed := flag.Bool("seed", false, "Run database seeder")
	migrateSeed := flag.Bool("migrate-seed", false, "Run migration and seeder")
	flag.Parse()

	// ===== LOAD ENV =====
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	// ===== CONNECT DB =====
	config.ConnectDB()

	// ===== HANDLE FLAGS =====
	if *migrateSeed {
		config.ResetDatabase()
		config.Migrate()
		seeders.SeedLevel()
		log.Println("Database reset, migration & seeder completed")
		return
	}

	if *migrate {
		config.Migrate()
		log.Println("Migration completed")
		return
	}

	if *seed {
		seeders.SeedLevel()
		log.Println("Seeder completed")
		return
	}

	// ===== NORMAL API MODE =====
	r := gin.Default()
	routes.Setup(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)
}
