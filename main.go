package main

import (
	"goapi/config"
	"goapi/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//DB
	config.ConnectDB()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:5173", "http://localhost:3000"}, //visit frontend
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders: []string{"Content-Length"},

		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routes.SetupRouters(r)

	r.Run(":8080")
}
