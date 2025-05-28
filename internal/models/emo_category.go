package models

type EmotionCategory struct {
	CategoryID uint   `gorm:"primaryKey;autoIncrement;column:category_id" json:"category_id"`
	Name       string `gorm:"type:varchar(50);not null" json:"name"`
}

func (EmotionCategory) TableName() string {
	return "emo_category"
}
