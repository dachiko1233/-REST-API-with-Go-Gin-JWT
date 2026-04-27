package routes

import (
	"goapi/handlers"
	"goapi/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouters(r *gin.Engine) {
	r.Use(middleware.Logger())

	r.POST("/api/v1/register", middleware.RateLimiter(), handlers.Register)
	r.POST("/api/v1/login", middleware.RateLimiter(), handlers.Login)
	r.GET("/api/v1/verify", handlers.VerifyEmail)
	r.POST("/api/v1/refresh", handlers.RefreshToken)

	api := r.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())

	{
		api.GET("/users", handlers.GetUsers)
		api.GET("/users/:id", handlers.GetUser)
		api.POST("/users", handlers.CreateUser)
		api.POST("/logout", handlers.Logout)
	}
}
