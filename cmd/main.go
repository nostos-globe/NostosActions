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
		ActionService: actionService,
		AuthClient:    authClient,
		ProfileClient: profileClient,
	}

	// Initialize Gin
	r := gin.Default()

	// Trip routes
	api := r.Group("/api/likes")
	{
		api.POST("/trip/:id", actionHandler.LikeTrip)
		api.DELETE("/trip/:id", actionHandler.UnlikeTrip)
		api.GET("/trip/:id", actionHandler.GetTripLikes)
	}

	followApi := r.Group("/api/actions")
	{
		followApi.POST("/create", actionHandler.CreateAction)
	}

	// Media routes in separate group
	favsApi := r.Group("/api/favourites")
	{
		favsApi.GET("/media/:id", actionHandler.GetMediaStatus)
		favsApi.POST("/media/:id", actionHandler.FavMedia)
		favsApi.DELETE("/media/:id", actionHandler.UnFavMedia)
	}

	// Start server
	log.Println("Server running on http://localhost:8086")
	if err := r.Run(":8086"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
