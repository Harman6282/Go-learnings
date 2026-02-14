// main.go

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Harman6282/auth-app/auth"
	"github.com/Harman6282/auth-app/db"
	"github.com/Harman6282/auth-app/handlers"
	"github.com/Harman6282/auth-app/middleware"
	"github.com/Harman6282/auth-app/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// loadEnv loads environment variables from .env file
func loadEnv() {
    // Load .env file if it exists
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using environment variables")
    }

    // Check required variables
    requiredVars := []string{"DATABASE_URL", "JWT_SECRET"}
    for _, v := range requiredVars {
        if os.Getenv(v) == "" {
            log.Fatalf("Required environment variable %s is not set", v)
        }
    }
}

func main() {
    // Load environment variables
    loadEnv()

    // Connect to the database
    database, err := db.Connect(os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    r := mux.NewRouter()

    // Create repositories
    userRepo := models.NewUserRepository(database)
    refreshTokenRepo := models.NewRefreshTokenRepository(database)

    // Create services
    authService := auth.NewAuthService(userRepo, refreshTokenRepo, os.Getenv("JWT_SECRET"), 15*time.Minute)

    // Create handlers
    authHandler := handlers.NewAuthHandler(authService)
    userHandler := handlers.NewUserHandler(userRepo)

    // Public routes
    r.HandleFunc("/api/auth/register", authHandler.Register).Methods("POST")
    r.HandleFunc("/api/auth/login", authHandler.Login).Methods("POST")
    r.HandleFunc("/api/auth/refresh", authHandler.RefreshToken).Methods("POST")

    // Protected routes
    protected := r.PathPrefix("/api").Subrouter()
    protected.Use(middleware.AuthMiddleware(authService))

    protected.HandleFunc("/profile", userHandler.Profile).Methods("GET")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Server starting on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}