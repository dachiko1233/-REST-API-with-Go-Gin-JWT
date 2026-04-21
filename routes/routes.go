package routes

import (
	"goapi/handlers"
	"goapi/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouters(r *gin.Engine) {
	r.Use(middleware.Logger())
	r.POST("/api/v1/login", handlers.Loggin)

	api := r.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())

	{
		api.GET("/users", handlers.GetUsers)
		api.GET("/users/:id", handlers.GetUser)
		api.POST("/users", handlers.CreateUser)
	}
}
