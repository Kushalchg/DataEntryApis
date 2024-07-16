package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string `gorm:"unique" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	CoformCode int    `json:"conform_code" gorm:"not null"`
	Verified   bool   `josn:"verified"`
	Updated    bool   `josn:"updated"`
}

type UserDetail struct {
	gorm.Model
	Email string `gorm:"unique" json:"email"`
	Name  string `gorm:"not null" json:"password"`
	Phone int    `json:"conform_code" gorm:"not null"`
}
