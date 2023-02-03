package model

import "gorm.io/gorm"

// User 用户模型结构
type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}

func (u *User) TableName() string {
	return "users"
}
