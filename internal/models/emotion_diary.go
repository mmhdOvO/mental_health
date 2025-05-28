package models

import "time"

type EmotionDiary struct {
	EmoID      uint            `gorm:"primaryKey;autoIncrement;column:emo_id" json:"emo_id"`
	UserID     uint            `gorm:"not null" json:"user_id"`    // 外键字段
	User       User            `gorm:"foreignKey:UserID" json:"-"` // 正确关联到 User 表
	CategoryID uint            `gorm:"not null" json:"category_id"`
	Category   EmotionCategory `gorm:"foreignKey:CategoryID" json:"-"` // 外键关联分类表
	EmoText    string          `gorm:"type:text;not null" json:"emo_text"`
	EmoStatus  string          `gorm:"type:varchar(20)" json:"emo_status"`
	EmoPhoto   string          `gorm:"type:varchar(255)" json:"emo_photo"`
	CreatedAt  time.Time       `gorm:"autoCreateTime" json:"created_at"`
}

func (EmotionDiary) TableName() string {
	return "emotion_diary"
}
