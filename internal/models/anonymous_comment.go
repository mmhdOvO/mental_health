package models

import "time"

type AnonymousComment struct {
	CommentID   uint          `gorm:"primaryKey;autoIncrement;column:comment_id" json:"comment_id"`
	PostID      uint          `gorm:"not null" json:"post_id"`
	Post        AnonymousPost `gorm:"foreignKey:PostID" json:"-"` // 外键关联帖子表
	AnonymousID string        `gorm:"type:varchar(36);not null" json:"anonymous_id"`
	Content     string        `gorm:"type:text;not null" json:"content"`
	CreatedAt   time.Time     `gorm:"autoCreateTime" json:"created_at"`
}

func (AnonymousComment) TableName() string {
	return "anonymous_comment"
}
