package main

import (
	"ceilo-backend/internal/config"
	"ceilo-backend/internal/database"
	"ceilo-backend/internal/middleware"
	"ceilo-backend/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	cfg := config.LoadConfig()

	// Initialize database connection
	db := database.InitDB(cfg)

	// Run migrations
	database.Migrate(db)

	// Seed initial data (optional)
	database.Seed(db)

	// Setup Gin router
	router := gin.Default()

	// Setup CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Setup routes
	routes.SetupRoutes(router, db)

	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}