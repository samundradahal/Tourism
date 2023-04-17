package model

import "gorm.io/gorm"

type Tourism struct {
	gorm.Model
	Name     string `json: "name"`
	Email    string `json:"email" binding:"required,email"`
	Language string `json:"language" binding:"required"`
}
