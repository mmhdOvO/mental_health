package models

import "time"

type Test struct {
	TestID   uint      `gorm:"primaryKey;autoIncrement;column:test_id" json:"test_id"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID" json:"-"` // 外键关联用户表
	TestType string    `gorm:"type:varchar(50);not null" json:"test_type"`
	Score    int       `gorm:"not null" json:"score"`
	TestTime time.Time `gorm:"autoCreateTime" json:"test_time"`
}

func (Test) TableName() string {
	return "psychological_test"
}
