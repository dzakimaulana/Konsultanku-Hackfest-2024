package main

import (
	"log"
	"os"

	"konsultanku-app/database"
	"konsultanku-app/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Environment is not detected")
	}

	cookieSecret := os.Getenv("COOKIE_SECRET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database.FirebaseConnection()
	database.DatabaseConnection()

	router := gin.Default()
	store := cookie.NewStore([]byte(cookieSecret))
	router.Use(sessions.Sessions("mysession", store))

	routes.AuthRouter(router)
	routes.StudentRouter(router)
	routes.TeamRouter(router)
	routes.MseRouter(router)

	// Jalankan aplikasi Anda
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
