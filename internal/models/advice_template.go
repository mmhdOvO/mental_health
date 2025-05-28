package models

type AdviceTemplate struct {
	TemplateID uint   `gorm:"primaryKey;autoIncrement;column:template_id" json:"template_id"`
	TestType   string `gorm:"type:varchar(50);not null" json:"test_type"`
	MinScore   int    `gorm:"not null" json:"min_score"`
	MaxScore   int    `gorm:"not null" json:"max_score"`
	AdviceText string `gorm:"type:text;not null" json:"advice_text"`
}

func (AdviceTemplate) TableName() string {
	return "advice_template"
}
