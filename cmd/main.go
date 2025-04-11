package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	controller "main/internal/api"
	dbRepo "main/internal/db"
	"main/internal/service"
	"main/pkg/config"
	"main/pkg/db"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or error loading it: %v", err)
	}

	secretsManager := config.GetSecretsManager()
	if secretsManager != nil {
		secrets := secretsManager.LoadSecrets()
		for key, value := range secrets {
			os.Setenv(key, value)
		}
	} else {
		log.Println("Falling back to environment variables")
	}
}

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Initialize repositories
	actionRepo := &dbRepo.ActionRepository{DB: database}

	// Initialize clients for external apis
	authClient := &service.AuthClient{BaseURL: cfg.AuthServiceUrl}
	profileClient := &service.ProfileClient{BaseURL: cfg.ProfileServiceUrl}

	// Initialize services
	actionService := &service.ActionService{ActionRepo: actionRepo}

	// Initialize controllers
	actionHandler := &controller.ActionController{
		ActionService:  actionService,
		AuthClient:    authClient,
		ProfileClient: profileClient,
	}

	// Initialize Gin
	r := gin.Default()

	// Trip routes
	api := r.Group("/api/likes")
	{
		api.POST("/trip/:id", albumHandler.LikeTrip)
		api.DELETE("/trip/:id", albumHandler.UnlikeTrip)
	}

	followApi := r.Group("/api/follows")
	{
		followApi.POST("/user/:id", albumHandler.FollowUser)
		followApi.DELETE("/user/:id", albumHandler.UnfollowUser)
	}

	uploadApi := r.Group("/api/uploads")
	{
		uploadApi.POST("/media", albumHandler.UploadMedia)
		uploadApi.DELETE("/media/:id", albumHandler.DeleteMedia)
		uploadApi.POST("/trip/:id", albumHandler.CreateTrip)
		uploadApi.DELETE("/trip/:id", albumHandler.DeleteTrip)
		uploadApi.POST("/album/:id", albumHandler.CreateAlbum)
		uploadApi.DELETE("/album/:id", albumHandler.DeleteAlbum)

	}
	// Media routes in separate group
	favsApi := r.Group("/api/favourites")
	{
		favsApi.POST("/media/:id", albumHandler.LikeMedia)
		favsApi.DELETE("/media/:id", albumHandler.UnlikeMedia)
	}

	// Start server
	log.Println("Server running on http://localhost:8086")
	if err := r.Run(":8086"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
