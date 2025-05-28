// models/user.go
package models

import (
	"time"
)

type User struct {
	UserID       uint      `gorm:"primaryKey;autoIncrement;column:user_id" json:"user_id"`
	Username     string    `gorm:"type:varchar(50);unique;not null" json:"username"` // 新增用户名字段
	Phone        string    `gorm:"type:varchar(15);unique;not null" json:"phone"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"` // 密码哈希字段（不返回前端）
	AvatarURL    string    `gorm:"type:varchar(255)" json:"avatar_url"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
