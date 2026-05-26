package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"mosquilab/config"
	"mosquilab/internal/database"
	"mosquilab/internal/handlers"
	"mosquilab/internal/middleware"
	"mosquilab/internal/repository"
)

func main() {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseURL)
	defer db.Close()

	// Repos
	eventRepo := repository.NewEventRepo(db)
	userRepo := repository.NewUserRepo(db)

	// Seed default admin (change password via DB!)
	defaultPass := os.Getenv("ADMIN_PASSWORD")
	if defaultPass == "" {
		defaultPass = "mosquilab2024"
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(defaultPass), bcrypt.DefaultCost)
	adminEmail := os.Getenv("ADMIN_EMAIL")
	if adminEmail == "" {
		adminEmail = "admin@mosquilab.univille.br"
	}
	if err := userRepo.SeedAdmin(adminEmail, string(hash)); err != nil {
		log.Printf("seed admin warning: %v", err)
	}

	// Handlers
	authH := handlers.NewAuthHandler(userRepo, cfg.JWTSecret)
	agendaH := handlers.NewAgendaHandler(eventRepo)

	// Router
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", cfg.CORSOrigin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		// Public
		api.GET("/agenda", agendaH.ListPublic)

		// Auth
		api.POST("/auth/login", authH.Login)

		// Admin (protected)
		admin := api.Group("/admin")
		admin.Use(middleware.JWTAuth(cfg.JWTSecret))
		{
			admin.GET("/agenda", agendaH.ListAll)
			admin.POST("/agenda", agendaH.Create)
			admin.PUT("/agenda/:id", agendaH.Update)
			admin.DELETE("/agenda/:id", agendaH.Delete)
		}
	}

	log.Printf("MosquiLab API starting on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
