package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model `swaggerignore:"true"`
	Username   string `gorm:"unique;not null" json:"username"`
	Password   string `gorm:"not null" json:"password"`
	IsAdmin       string `gorm:"not null" json:"isAdmin"`
	Nif        string `gorm:"not null" json:"nif"`
}
