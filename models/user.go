package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `jsom:"email"`
	Age   int    `json:"age"`
}

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
