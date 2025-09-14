package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string
}

// UserResponse 用户响应模型（不包含敏感信息）
type UserResponse struct {
	gorm.Model
	Username string
	Email    string
}
