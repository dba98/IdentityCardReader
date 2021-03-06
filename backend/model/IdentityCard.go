package model

import (
	"github.com/jinzhu/gorm"
)

type IdentityCard struct {
	gorm.Model `swaggerignore:"true"`
	Nif        string `gorm:"unique;" json:"nif"`
	Image1     string `gorm:"unique;" json:"frontData"`
	Image2     string `gorm:"unique;" json:"backData"`
}
