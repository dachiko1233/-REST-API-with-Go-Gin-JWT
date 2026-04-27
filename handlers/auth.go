package handlers

import (
	"fmt"
	"goapi/config"
	"goapi/models"
	"goapi/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("Dachi1234")

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

	//generate verification token

	token := fmt.Sprintf("%d", time.Now().UnixNano())

	user := models.User{
		Name:              req.Name,
		Email:             req.Email,
		Password:          string(hashedPassword),
		Age:               req.Age,
		IsVerified:        false,
		VerificationToken: token,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	//send verification email

	if err := utils.SendVerificationEmail(user.Email, token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not send verification email"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful! Please check your email to verify your account.",
	})

	// c.JSON(http.StatusCreated, models.UserResponse{
	// 	ID:    user.ID,
	// 	Name:  user.Name,
	// 	Email: user.Email,
	// })

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
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	if !user.IsVerified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please verify your email"})
		return
	}

	//Generate Refresh
	refreshTokenString := fmt.Sprintf("%d-%d", user.ID, time.Now().UnixNano())

	refreshToken := models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshTokenString,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	config.DB.Create(&refreshToken)
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "refresh_token": refreshTokenString, "expires_in": 900})

}

func VerifyEmail(c *gin.Context) {
	token := c.Query("token")

	var user models.User

	if err := config.DB.Where("verification_token = ?", token).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid verification token"})
		return
	}

	config.DB.Model(&user).Updates(map[string]interface{}{
		"is_verified":        true,
		"verification_token": "",
	})

	c.JSON(http.StatusOK, gin.H{"message": "Email verified seccessfully! You can now login"})

}

func Logout(c *gin.Context) {
	//get token from handler

	authHeader := c.GetHeader("Authorization")
	parts := strings.Split(authHeader, " ")
	tokenString := parts[1]

	//save token to blacklist
	blacklisted := models.BlacklistedToken{
		Token: tokenString,
	}

	if err := config.DB.Create(&blacklisted).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not lngout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully!"})
}

func RefreshToken(c *gin.Context) {
	var body struct {
		ResponseToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//find refresh token in DB
	var refreshToken models.RefreshToken
	if err := config.DB.Where("token = ?", body.ResponseToken).First(&refreshToken).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	//Check if expered
	if time.Now().Unix() > refreshToken.ExpiresAt {
		config.DB.Delete(&refreshToken)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token expired, please login again"})
		return
	}

	//find user
	var user models.User
	config.DB.First(&user, refreshToken.UserID)
	//Generate new Access tokken
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	})

	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessTokenString,
		"expires_in":   900,
	})

}
