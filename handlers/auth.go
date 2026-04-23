package handlers

import (
	"goapi/config"
	"goapi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("Dachi1234")

var fakeUsers = map[string]string{
	"alice@example.com": "password123@",
	"bob@example.com":   "password456",
}

func Loggin(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//if users password mach

	pass, exists := fakeUsers[body.Email]
	if !exists || pass != body.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentals"})
		return
	}

	//generate JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": body.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Register(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//has the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Age:      req.Age,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	c.JSON(http.StatusCreated, models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})

}

func Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// check password
	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credenttials"})
		return
	}

	//generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

}
