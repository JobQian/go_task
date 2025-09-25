package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string
	Content string
	UserId  int
	User    User `gorm:"foreignKey:UserId"`
}
