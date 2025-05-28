package models

import "time"

type AnonymousPost struct {
	PostID      uint      `gorm:"primaryKey;autoIncrement;column:post_id" json:"post_id"`
	AnonymousID string    `gorm:"type:varchar(36);not null" json:"anonymous_id"` // 使用UUID
	Content     string    `gorm:"type:text;not null" json:"content"`
	IsHidden    bool      `gorm:"default:false" json:"is_hidden"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (AnonymousPost) TableName() string {
	return "anonymous_post"
}
