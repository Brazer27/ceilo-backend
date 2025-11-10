package routes

import (
	"ceilo-backend/internal/handler"
	"ceilo-backend/internal/middleware"
	"ceilo-backend/internal/repository"
	"ceilo-backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	forumRepo := repository.NewForumRepository(db)
	consultationRepo := repository.NewConsultationRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)
	forumService := service.NewForumService(forumRepo)
	consultationService := service.NewConsultationService(consultationRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(userService)
	forumHandler := handler.NewForumHandler(forumService)
	consultationHandler := handler.NewConsultationHandler(consultationService)

	// Health check endpoint
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"message": "Ceilo API is running",
		})
	})

	// API routes
	api := router.Group("/api")
	{
		// Authentication routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes (require JWT)
		protected := api.Group("")
		protected.Use(middleware.JWTAuthMiddleware())
		{
			// Profile routes
			protected.GET("/profile", authHandler.GetProfile)

			// Forum routes
			forum := protected.Group("/forum")
			{
				forum.POST("", forumHandler.CreateForum)
				forum.GET("", forumHandler.GetAllForums)
				forum.GET("/:id", forumHandler.GetForumByID)
				forum.PUT("/:id", forumHandler.UpdateForum)
				forum.DELETE("/:id", forumHandler.DeleteForum)
			}

			// Consultation routes
			consultation := protected.Group("/consultation")
			{
				consultation.POST("", consultationHandler.CreateConsultation)
				consultation.GET("", consultationHandler.GetUserConsultations)
				consultation.GET("/:id", consultationHandler.GetConsultationByID)
				consultation.PUT("/:id/status", consultationHandler.UpdateConsultationStatus)
			}

			// Stress test routes (placeholder)
			protected.GET("/stress-test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Stress test endpoint"})
			})

			// Articles routes (placeholder)
			protected.GET("/articles", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Articles endpoint"})
			})

			// Events routes (placeholder)
			protected.GET("/events", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Events endpoint"})
			})
		}
	}
}