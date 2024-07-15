package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// CoformCode int    `json:"conform_code"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
