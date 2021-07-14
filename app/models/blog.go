package models

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	CustomModel

	HeaderImage string
	Title       string
	Content     string `gorm:"type:text"`

	UserID uint
	User   User
}
