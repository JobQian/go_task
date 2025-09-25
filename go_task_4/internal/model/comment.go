package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	UserId  int
	User    User `gorm:"foreignKey:UserId"`
	PostId  int
	Post    Post `gorm:"foreignKey:PostId"`
}
