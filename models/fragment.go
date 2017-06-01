package models

import "github.com/jinzhu/gorm"

type Fragment struct {
	gorm.Model
	Contents string `json:"contents" form:"contents"`
}
