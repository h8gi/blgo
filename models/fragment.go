package models

import "github.com/jinzhu/gorm"

type Fragment struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"not null;unique"`
	Contents string `json:"contents" form:"contents" gorm:"not null"`
}
