package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sajid414/Go-react-auth/database"
	"github.com/sajid414/Go-react-auth/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("error in loading env file")
	}
	database.ConnectDB()
}

func main() {
	sqlDB, err := database.DBConn.DB()
	if err != nil {
		log.Println("Error in database Connection.")
	}
	defer sqlDB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// ✅ CORS middleware setup
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetupRoutes(router)

	log.Fatal(router.Run(":" + port))
}
