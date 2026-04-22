package main

import (
	"goapi/config"
	"goapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	//DB
	config.ConnectDB()

	r := gin.Default()

	routes.SetupRouters(r)

	r.Run(":8080")
}
