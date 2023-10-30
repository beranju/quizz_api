package models

import (
	"github.com/jinzhu/gorm"
)

type Quiz struct {
	gorm.Model
	QuizName    string     `json:"quiz_name" gorm:"quiz_name;not null"`
	Description string     `json:"description" gorm:"description;not null"`
	Duration    int64      `json:"duration" gorm:"duration;not null"`
	Question    []Question `gorm:"foreignKey:quiz_id"`
}
