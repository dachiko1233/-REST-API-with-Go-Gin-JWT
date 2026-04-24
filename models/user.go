package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name              string `json:"name"`
	Email             string `json:"email"`
	Password          string `json:"-"`
	Age               int    `json:"age"`
	IsVerified        bool   `json:"is_verifed" gorm:"default:false"`
	VerificationToken string `json:"-"`
}

type CreateUserRequest struct {
	Name     string `json:"name"  binding:"required,min=2,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binfing:"required,min=6"`
	Age      int    `json:"age"   binding:"required,min=1,max=120"`
}

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BlacklistedToken struct {
	gorm.Model
	Token string `gorm:"unique"`
}
