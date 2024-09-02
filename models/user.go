package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type UpdateProfileInput struct {
	Id       uint   `json:"user_id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}
