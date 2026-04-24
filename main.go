package main

import (
	"goapi/config"
	"goapi/routes"
	"log"

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
	routes.SetupRouters(r)
	r.Run(":8080")
}
