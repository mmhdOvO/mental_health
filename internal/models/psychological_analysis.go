package models

import (
	"gorm.io/datatypes"
	"time"
)

type PsychologicalAnalysis struct {
	AnalyseID  uint           `gorm:"primaryKey;autoIncrement;column:analyse_id" json:"analyse_id"`
	UserID     uint           `gorm:"not null" json:"user_id"`
	User       User           `gorm:"foreignKey:UserID" json:"-"` // 外键关联用户表
	RiskLevel  string         `gorm:"type:varchar(20)" json:"risk_level"`
	AIReport   datatypes.JSON `gorm:"type:json" json:"ai_report"` // 使用GORM的JSON类型
	AnalysedAt time.Time      `gorm:"autoCreateTime" json:"analysed_at"`
}

func (PsychologicalAnalysis) TableName() string {
	return "psychological_analysis"
}
