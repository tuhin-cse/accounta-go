package main

import (
	"accounta-go/config"
	"accounta-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var err error

func init() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to database!")
	}
	config.MigrateDB()
	r := gin.Default()
	routes.ApiRoutes(r)
	err = r.Run(":" + port)
	if err != nil {
		return
	}
}
