package database

import (
	"ceilo-backend/internal/models"
	"log"

	"gorm.io/gorm"
)

// Migrate runs all database migrations
func Migrate(db *gorm.DB) {
	log.Println("Running database migrations...")

	err := db.AutoMigrate(
		&models.User{},
		&models.Forum{},
		&models.ForumComment{},
		&models.StressTest{},
		&models.Consultation{},
		&models.Article{},
		&models.Event{},
	)

	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	log.Println("Database migrations completed successfully")
}