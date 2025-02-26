package main

import (
	"fmt"
	"log"
	"os"
	"mentalartsapi/handlers"
	"mentalartsapi/models"
	_ "mentalartsapi/docs" // swagger docs

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           Book Library API
// @version         1.0
// @description     A Book Library Management System API
// @termsOfService  http://swagger.io/terms/

// @contact.name   Anıl Yağız
// @contact.email  a.yagiz@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8000
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it.")
	}

	// Set default values if environment variables are not set
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "123abcd")
	dbName := getEnv("DB_NAME", "postgres")
	dbPort := getEnv("DB_PORT", "5432")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")
	apiPort := getEnv("API_PORT", "8000")

	// Configure Gin mode
	ginMode := getEnv("GIN_MODE", "debug")
	gin.SetMode(ginMode)

	// Connect to database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect database: %v", err)
	}

	// Auto migrate models
	db.AutoMigrate(&models.Author{}, &models.Book{}, &models.Review{})

	// Initialize DB in handlers
	handlers.InitDB(db)

	// Create router
	router := gin.Default()

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Authors routes
		v1.POST("/authors", handlers.CreateAuthor)
		v1.GET("/authors", handlers.GetAllAuthors)
		v1.GET("/authors/:id", handlers.GetAuthor)
		v1.PUT("/authors/:id", handlers.UpdateAuthor)
		v1.DELETE("/authors/:id", handlers.DeleteAuthor)

		// Books routes
		v1.POST("/books", handlers.CreateBook)
		v1.GET("/books", handlers.GetAllBooks)
		v1.GET("/books/:id", handlers.GetBook)
		v1.PUT("/books/:id", handlers.UpdateBook)
		v1.DELETE("/books/:id", handlers.DeleteBook)

		// Reviews routes
		v1.GET("/books/:id/reviews", handlers.GetBookReviews)
		v1.POST("/books/:id/reviews", handlers.CreateReview)
		v1.PUT("/reviews/:id", handlers.UpdateReview)
		v1.DELETE("/reviews/:id", handlers.DeleteReview)
	}

	// Test routes
	router.GET("/ping", handlers.HandlePing)
	router.GET("/hello", handlers.HandleHello)
	router.GET("/helloWithPayload", handlers.HandleHelloWithPayload)

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server
	log.Printf("Server starting on port %s...\n", apiPort)
	router.Run(fmt.Sprintf(":%s", apiPort))
}

// getEnv gets value from environment or returns default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Vanilla implementation
// func main() {
// 	http.HandleFunc("GET /ping", handlePing)
// 	log.Println("Server listening...")
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

// func handlePing(w http.ResponseWriter, r *http.Request) {
// 	res := Response{Msg: "pong"}
// 	json.NewEncoder(w).Encode(res)
// 	w.WriteHeader(http.StatusOK)
// 	log.Println("Request recieved")
// }
