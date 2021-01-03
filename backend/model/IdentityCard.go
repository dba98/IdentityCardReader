package model

import "github.com/jinzhu/gorm"

type IdentityCard struct {
	gorm.Model `swaggerignore:"true"`
	Name   string `gorm:"not null" json:"name"`
	NumberId   string `gorm:"not null" json:"numberid"`
	Nif      string `gorm: "unique;not null" json:"nif"`
	BirthDay string `gorm:"not null"`
}
