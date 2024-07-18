package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email      string      `gorm:"unique" json:"email"`
	Password   string      `gorm:"not null" json:"password"`
	CoformCode int         `json:"conform_code" gorm:"not null"`
	Verified   bool        `josn:"verified"`
	Updated    bool        `josn:"updated"`
	Entries    []EntryData `gorm:"foreignKey:UId"`
}

type UserDetail struct {
	gorm.Model
	Email     string    `gorm:"unique" json:"email"`
	Name      string    `gorm:"not null" json:"password"`
	Phone     int       `json:"conform_code" gorm:"not null"`
	EntryData EntryData `gorm:"references:ID"`
}

type EntryData struct {
	gorm.Model
	Tname     string  `gorm:"not null" json:"tname"`
	Tlength   float32 `gorm:"not null" json:"tlength"`
	Tdiameter float32 `josn:"tdiameter"`
	Tlogi     float32 `josn:"tlongi"`
	Tlatt     float32 `josn:"tlatt"`
	UId       int     `josn:"uid" gorm:"not null"`
}
