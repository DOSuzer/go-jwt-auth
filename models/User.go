package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	Email    string `gorm:"unique" json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role"`
}

type UserNameChange struct {
	Name string `json:"name" binding:"required"`
}
