package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserResponse 用户响应模型（不包含敏感信息）
type UserResponse struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
}
