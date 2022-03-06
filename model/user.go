package model

import (
	_ "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not null" json:"username"`
	Email    string `gorm:"unique_index;not null" json:"email" validate:"required,email,min=6,max=32"`
	Password string `gorm:"not null" json:"password" validate:"required,min=6"`
	Names    string `json:"names"`
}
