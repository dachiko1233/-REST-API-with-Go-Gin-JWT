package handlers

import (
	"fmt"
	"goapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 30},
	{ID: 2, Name: "Bob", Email: "bob@example.com", Age: 25},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Received ID:", id)
	fmt.Printf("ID type checkd %s", id)

	for _, user := range users {
		if fmt.Sprintf("%d", user.ID) == id {
			c.JSON(http.StatusOK, gin.H{"data": user})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
}

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newUser.ID = uint(len(users) + 1)
	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{"data": newUser})
}
