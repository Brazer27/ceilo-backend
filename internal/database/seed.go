package database

import (
	"ceilo-backend/internal/models"
	"ceilo-backend/internal/utils"
	"log"
	"time"

	"gorm.io/gorm"
)

// Seed populates database with initial data
func Seed(db *gorm.DB) {
	log.Println("Seeding database...")

	// Check if admin user already exists
	var count int64
	db.Model(&models.User{}).Where("email = ?", "admin@ceilo.com").Count(&count)
	
	if count > 0 {
		log.Println("Database already seeded, skipping...")
		return
	}

	// Create admin user
	hashedPassword, _ := utils.HashPassword("admin123")
	admin := models.User{
		Name:     "Admin Ceilo",
		Email:    "admin@ceilo.com",
		Password: hashedPassword,
		Role:     "admin",
	}
	db.Create(&admin)

	// Create sample articles
	articles := []models.Article{
		{
			Title:       "Mengenal Kesehatan Mental",
			Content:     "Kesehatan mental adalah aspek penting dalam kehidupan...",
			Category:    "Edukasi",
			AuthorID:    admin.ID,
			PublishedAt: time.Now(),
		},
		{
			Title:       "Tips Mengelola Stres",
			Content:     "Stres adalah bagian dari kehidupan yang tidak dapat dihindari...",
			Category:    "Tips",
			AuthorID:    admin.ID,
			PublishedAt: time.Now(),
		},
	}
	db.Create(&articles)

	log.Println("Database seeding completed successfully")
}