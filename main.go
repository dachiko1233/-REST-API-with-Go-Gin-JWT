package main

import (
	"goapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRouters(r)

	r.Run(":8080")
}
