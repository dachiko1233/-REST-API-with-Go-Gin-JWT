package handlers

import (
	"goapi/config"
	"goapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
	var users []models.User

	id := c.Param("id")

	if err := config.DB.First(&users, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	config.DB.Create(&user)

	c.JSON(http.StatusCreated, models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
	})
}
