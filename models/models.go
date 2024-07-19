package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string      `gorm:"unique" json:"email"`
	Password   string      `gorm:"not null" json:"password"`
	CoformCode int         `json:"conform_code" gorm:"not null"`
	Verified   bool        `json:"verified"`
	Updated    bool        `json:"updated"`
	Entries    []EntryData `json:"entries" gorm:"foreignKey:UId;references:ID"`
}

type UserDetail struct {
	gorm.Model
	Email string `gorm:"unique" json:"email"`
	Name  string `gorm:"not null" json:"password"`
	Phone int    `json:"conform_code" gorm:"not null"`
}

type EntryData struct {
	gorm.Model
	Tname     string  `gorm:"not null" json:"tname"`
	Tlength   float32 `gorm:"not null" json:"tlength"`
	Tdiameter float32 `json:"tdiameter"`
	Tlongi    float32 `json:"tlongi"`
	Tlatt     float32 `json:"tlatt"`
	UId       int     `json:"uid" gorm:"not null"`
}
