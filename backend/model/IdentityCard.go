package model

import (
	"github.com/jinzhu/gorm"
)

type IdentityCard struct {
	gorm.Model `swaggerignore:"true"`
	Name       string `gorm:"not null" json:"name"`
	NumberID   string `gorm:"not null" json:"numberid"`
	Nif        string `gorm:"unique;not null" json:"nif"`
	Birthday   string `gorm:"not null" json:"birthday"`
	Image1     string `gorm:"unique;not null" json:"image1"`
	Image2     string `gorm:"unique;not null" json:"image2"`
}
