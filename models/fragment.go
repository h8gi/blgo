package models

import "github.com/jinzhu/gorm"

type Fragment struct {
	gorm.Model
	Text string `json:"text" form:"text"`
}
